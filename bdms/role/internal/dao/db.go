package dao

import (
	"context"
	dsql "database/sql"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	pb "role/api"
	"role/internal/model"
	con "role/tool/container"
)

const (
	//顶级菜单Id
	_topMenuId = 1

	//表名
	//菜单表
	_sysRoleTable = "sys_role"

	_sysRoleMenuTable = "sys_role_menu"

	//SQL语句
	//查询角色列表
	_queryRoleList = "select id, num, role_name, role_desc from %s order by num"

	//查询角色菜单列表
	_queryRoleMenuList = "select menu_id from %s where role_id = ?"

	//插入角色
	_insertRole = "insert into %s (num, role_name, role_desc) values (?,?,?)"

	//更新角色
	_updateRole = "update %s set num = ?, role_name = ?, role_desc = ? where id = ?"

	//删除角色
	_deleteRole = "delete from %s where id = ?"

	//删除角色菜单(按角色id)
	_deleteRoleMenus = "delete from %s where role_id = ?"

	//插入角色菜单
	_insertRoleMenu = "insert into %s (role_id, menu_id) values (?,?)"

	//删除某角色菜单(按角色id和菜单id)
	_deleteRoleMenu = "delete from %s where role_id = ? and menu_id = ?"

	//删除某些角色菜单(按菜单id)
	_deleteRoleNullMenus = "delete from %s where menu_id = ?"

	//查询角色菜单列表
	_queryRoleOptions = "select id, role_name from %s"
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

//查询角色列表
func (d *dao) RawRoleList(ctx context.Context) (resp *pb.GetRoleListResp, err error) {
	sqlRole := fmt.Sprintf(_queryRoleList, _sysRoleTable)
	sqlRoleMenu := fmt.Sprintf(_queryRoleMenuList, _sysRoleMenuTable)
	rows, err := d.db.Query(ctx, sqlRole)
	if err != nil {
		log.Error("select role err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.GetRoleListResp{}
	for rows.Next() {
		role := &pb.GetRoleListResp_Role{}
		if err = rows.Scan(&role.Id, &role.Num, &role.RoleName, &role.RoleDesc); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_role]] scan all record error(%v)", err)
			return
		}
		rowsRoleMenu, err1 := d.db.Query(ctx, sqlRoleMenu, role.Id)
		if err1 != nil {
			log.Error("select role_menu err(%v)", err1)
			err1 = err
			return
		}
		menusReq := &pb.GetMenusReq{}
		menusReq.RoleId = role.Id
		var menusResp *pb.GetMenusResp
		if menusResp, err = d.menuClient.GetMenus(ctx, menusReq); err != nil {
			log.Error("dao.menuClient.GetMenus err(%v)", err)
			return
		}
		role.Children = menusResp.Menus
		resp.RoleList = append(resp.RoleList, role)
		rowsRoleMenu.Close()
	}
	return
}

//插入角色
func (d *dao) InsertRole(ctx context.Context, req *pb.AddRoleReq) (resp *pb.AddRoleResp, err error) {
	sql := fmt.Sprintf(_insertRole, _sysRoleTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sql, req.Num, req.RoleName, req.RoleDesc); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_role]] failed to insert error(%v)", err)
		return
	}
	resp = &pb.AddRoleResp{}
	var roleId int64
	if roleId, err = res.LastInsertId(); err != nil {
		resp.Result = "error"
		return
	}
	setRoleRightsReq := &pb.SetRoleRightsReq{}
	setRoleRightsReq.RoleId = roleId
	setRoleRightsResp, err := d.SetRoleMenus(ctx, setRoleRightsReq)
	if err != nil {
		log.Error("dao.SetRoleMenus err(%v)", err)
		return
	}
	resp.Result = setRoleRightsResp.Result
	return
}

//更新角色
func (d *dao) UpdateRole(ctx context.Context, req *pb.UpdateRoleReq) (resp *pb.UpdateRoleResp, err error) {
	sql := fmt.Sprintf(_updateRole, _sysRoleTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sql, req.Num, req.RoleName, req.RoleDesc, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_role]] failed to update: (%v), error(%v)", req.Id, err)
	}
	resp = &pb.UpdateRoleResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//删除角色
func (d *dao) DeleteRole(ctx context.Context, req *pb.DeleteRoleReq) (resp *pb.DeleteRoleResp, err error) {
	sqlRole := fmt.Sprintf(_deleteRole, _sysRoleTable)
	sqlRoleMenu := fmt.Sprintf(_deleteRoleMenus, _sysRoleMenuTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlRoleMenu, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_role_menu]] failed to delete: (%v), error(%v)", req.Id, err)
		return
	}
	if res, err = d.db.Exec(ctx, sqlRole, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_role]] failed to delete: (%v), error(%v)", req.Id, err)
		return
	}
	resp = &pb.DeleteRoleResp{}
	var row int64
	if row, err = res.RowsAffected(); err == nil && row >= 1 {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//查询角色菜单
func (d *dao) RawRoleMenus(ctx context.Context, req *pb.GetRoleRightsReq) (resp *pb.GetRoleRightsResp, err error) {
	sql := fmt.Sprintf(_queryRoleMenuList, _sysRoleMenuTable)
	rows, err := d.db.Query(ctx, sql, req.RoleId)
	if err != nil {
		log.Error("select sys_role_menu err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.GetRoleRightsResp{}
	for rows.Next() {
		var roleId int64
		if err = rows.Scan(&roleId); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_role_menu]] scan all record error(%v)", err)
			return
		}
		resp.MenusId = append(resp.MenusId, roleId)
	}
	return
}

//设置角色菜单
func (d *dao) SetRoleMenus(ctx context.Context, req *pb.SetRoleRightsReq) (resp *pb.SetRoleRightsResp, err error) {
	sqlInsertRoleMenu := fmt.Sprintf(_insertRoleMenu, _sysRoleMenuTable)
	sqlDeleteRoleMenu := fmt.Sprintf(_deleteRoleMenu, _sysRoleMenuTable)
	roleRightReq := &pb.GetRoleRightsReq{}
	roleRightReq.RoleId = req.RoleId
	roleRightResp, err := d.RawRoleMenus(ctx, roleRightReq)
	if err != nil {
		log.Error("dao.GetRoleMenu err(%v)", err)
		return
	}
	newMenusId := con.New(req.MenusId...)
	oldMenusId := con.New(roleRightResp.MenusId...)
	newMenusId.Add(_topMenuId)
	inertMenusId := newMenusId.Minus(oldMenusId).List()
	deleteMenusId := oldMenusId.Minus(newMenusId).List()
	var res dsql.Result
	for _, v := range inertMenusId {
		if res, err = d.db.Exec(ctx, sqlInsertRoleMenu, req.RoleId, v); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_role_menu]] failed to insert error(%v)", err)
			return
		}
	}
	for _, v := range deleteMenusId {
		if res, err = d.db.Exec(ctx, sqlDeleteRoleMenu, req.RoleId, v); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_role_menu]] failed to insert error(%v)", err)
			return
		}
	}
	resp = &pb.SetRoleRightsResp{}
	if len(inertMenusId) <= 0 && len(deleteMenusId) <= 0 {
		resp.Result = "success"
		return
	}
	if _, err = res.LastInsertId(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//删除角色菜单
func (d *dao) DeleteRoleMenus(ctx context.Context, req *pb.DeleteRoleRightsReq) (resp *pb.DeleteRoleRightsResp, err error) {
	sql := fmt.Sprintf(_deleteRoleMenu, _sysRoleMenuTable)
	childrenMenuListReq := &pb.GetChildrenMenuListReq{}
	childrenMenuListReq.MenuId = req.MenuId
	var childrenMenuListResp *pb.GetChildrenMenuListResp
	if childrenMenuListResp, err = d.menuClient.GetChildrenMenuList(ctx, childrenMenuListReq); err != nil {
		log.Error("dao.menuClient.GetChildrenMenuList err(%v)", err)
		return
	}
	childrenMenuListResp.MenusId = append(childrenMenuListResp.MenusId, req.MenuId)
	var res dsql.Result
	for _, v := range childrenMenuListResp.MenusId {
		if res, err = d.db.Exec(ctx, sql, req.RoleId, v); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_role_menu]] failed to delete: (%v), (%v), error(%v)", req.RoleId, v, err)
			return
		}
	}
	resp = &pb.DeleteRoleRightsResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//删除角色菜单不存在的菜单
func (d *dao) DeleteRoleNullMenus(ctx context.Context, req *pb.DeleteRoleNullRightsReq) (resp *pb.DeleteRoleNullRightsResp, err error) {
	sql := fmt.Sprintf(_deleteRoleNullMenus, _sysRoleMenuTable)
	var res dsql.Result
	for _, v := range req.MenusId {
		if res, err = d.db.Exec(ctx, sql, v); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_role_menu]] failed to delete: (%v), error(%v)", v, err)
			return
		}
	}
	resp = &pb.DeleteRoleNullRightsResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//查询菜单选项
func (d *dao) RawRoleOptions(ctx context.Context) (resp *pb.GetRoleOptionsResp, err error) {
	sql := fmt.Sprintf(_queryRoleOptions, _sysRoleTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select role err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.GetRoleOptionsResp{}
	for rows.Next() {
		roleOption := &pb.GetRoleOptionsResp_RoleOption{}
		if err = rows.Scan(&roleOption.Id, &roleOption.RoleName); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_role]] scan all record error(%v)", err)
			return
		}
		resp.RoleOptions = append(resp.RoleOptions, roleOption)
	}
	return
}
