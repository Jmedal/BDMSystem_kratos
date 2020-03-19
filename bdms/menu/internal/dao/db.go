package dao

import (
	"context"
	dsql "database/sql"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	pb "menu/api"
	"menu/internal/model"
)

const (
	//表名
	//菜单表
	_sysMenuTable = "sys_menu"

	//SQL语句
	//查询菜单列表
	_queryMenuList = "select id, pid, num, menu_name, icon, path, levels from %s where pid = ? and levels = ? order by num"

	//插入菜单
	_insertMenu = "insert into %s (pid, num, menu_name, icon, path, levels) VALUES (?,?,?,?,?,?)"

	//查询菜单层级
	_queryMenuLevels = "select levels from %s where id = ?"

	//更新菜单层级
	_updateMenuLevels = "update %s set levels = levels + ? WHERE id = ?"

	//更新菜单
	_updateMenu = "update %s set pid = ?, num = ?, menu_name = ?, icon = ?, path = ?, levels = ? WHERE id = ?"

	//删除菜单
	_deleteMenu = "delete from %s where id = ?"

	//查询菜单选项
	_queryMenuOption = "select id, menu_name from %s where pid = ? and levels = ? order by num"

	//查询菜单
	//_queryMenus = "select id, pid, num, menu_name, icon, path, levels from %s where id = ?"

)

func NewDB() (db *sql.DB, cf func(), err error) {
	var (
		cfg sql.Config
		ct  paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Mysql").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db = sql.NewMySQL(&cfg)
	cf = func() { db.Close() }
	return
}

func (d *dao) RawArticle(ctx context.Context, id int64) (art *model.Article, err error) {
	// get data from db
	return
}

//查询菜单列表
func (d *dao) RawMenuList(ctx context.Context, pid int64, levels int64) (resp *pb.GetMenuListResp, err error) {
	sql := fmt.Sprintf(_queryMenuList, _sysMenuTable)
	rows, err := d.db.Query(ctx, sql, pid, levels+1)
	if err != nil {
		log.Error("select menu err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.GetMenuListResp{}
	for rows.Next() {
		menu := &pb.GetMenuListResp_Menu{}
		if err = rows.Scan(&menu.Id, &menu.Pid, &menu.Num, &menu.MenuName, &menu.Icon, &menu.Path, &menu.Levels); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_menu]] scan all record error(%v)", err)
			return
		}
		resp.MenuList = append(resp.MenuList, menu)
		children := &pb.GetMenuListResp{}
		if children, err = d.RawMenuList(ctx, menu.Id, menu.Levels); err != nil {
			log.Error("recursion err(%v)", err)
			return
		}
		resp.MenuList = append(resp.MenuList, children.MenuList...)
	}
	return
}

//添加菜单
func (d *dao) InsertMenu(ctx context.Context, req *pb.AddMenuReq) (resp *pb.AddMenuResp, err error) {
	sql := fmt.Sprintf(_insertMenu, _sysMenuTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sql, req.Pid, req.Num, req.MenuName, req.Icon, req.Path, req.Levels); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_menu]] failed to insert error(%v)", err)
		return
	}
	resp = &pb.AddMenuResp{}
	if _, err = res.LastInsertId(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//更新菜单
func (d *dao) UpdateMenu(ctx context.Context, req *pb.UpdateMenuReq) (resp *pb.UpdateMenuResp, err error) {
	sqlQueryLevels := fmt.Sprintf(_queryMenuLevels, _sysMenuTable)
	sqlUpdateLevels := fmt.Sprintf(_updateMenuLevels, _sysMenuTable)
	sqlUpdate := fmt.Sprintf(_updateMenu, _sysMenuTable)
	var oldLevels int64
	if err = d.db.QueryRow(ctx, sqlQueryLevels, req.Id).Scan(&oldLevels); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_menu]] scan short id record error(%v)", err)
		return
	}
	oldLevels = req.Levels - oldLevels
	if oldLevels != 0 {
		for i := 0; i < len(req.ChildrenId); i++ {
			if _, err = d.db.Exec(ctx, sqlUpdateLevels, oldLevels, req.ChildrenId[i]); err != nil {
				log.Error("[dao.dao-anchor.mysql|db[sys_menu]] failed to update: (%v), error(%v)", req.ChildrenId[i], err)
			}
		}
	}
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlUpdate, req.Pid, req.Num, req.MenuName, req.Icon, req.Path, req.Levels, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_menu]] failed to update: (%v), error(%v)", req.Id, err)
	}
	resp = &pb.UpdateMenuResp{}
	if _, err = res.RowsAffected(); err == nil{
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//删除菜单
func (d *dao) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) (resp *pb.DeleteMenuResp, err error) {
	sql := fmt.Sprintf(_deleteMenu, _sysMenuTable)
	for i := 0; i < len(req.ChildrenId); i++ {
		if _, err = d.db.Exec(ctx, sql, req.ChildrenId[i]); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_menu]] failed to delete: (%v), error(%v)", req.ChildrenId[i], err)
		}
	}
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sql, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_menu]] failed to delete: (%v), error(%v)", req.Id, err)
		return
	}
	resp = &pb.DeleteMenuResp{}
	var row int64
	if row, err = res.RowsAffected(); err == nil && row >= 1 {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//查询菜单选项
func (d *dao) RawMenuOptions(ctx context.Context, pid int64, levels int64, req *pb.GetMenuOptionsReq) (resp *pb.GetMenuOptionsResp, err error) {
	if levels > req.MinLevels-1 {
		return nil, nil
	}
	sql := fmt.Sprintf(_queryMenuOption, _sysMenuTable)
	rows, err := d.db.Query(ctx, sql, pid, levels)
	if err != nil {
		log.Error("select menu err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.GetMenuOptionsResp{}
	for i := 0; rows.Next(); i++ {
		menu := &pb.GetMenuOptionsResp_MenuOption{}
		if err = rows.Scan(&menu.Id, &menu.MenuName); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_menu]] scan all record error(%v)", err)
			return
		}
		children := &pb.GetMenuOptionsResp{}
		if children, err = d.RawMenuOptions(ctx, menu.Id, levels+1, req); err != nil {
			log.Error("recursion err(%v)", err)
			return
		}
		if children == nil {
			menu.Children = nil
		} else {
			if len(children.MenusOptions) <= 0 {
				continue
			}
			menu.Children = children.MenusOptions
		}
		resp.MenusOptions = append(resp.MenusOptions, menu)
	}
	return
}
