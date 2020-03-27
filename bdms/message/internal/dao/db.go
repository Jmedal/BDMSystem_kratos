package dao

import (
	"context"
	dsql "database/sql"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
	pb "message/api"
	"message/internal/model"
	"time"
)

const (
	//未读
	_messageUnRead = 1

	//已读
	_messageRead = 2

	//表
	_sysMessageTable = "sys_message"

	_sysMessageUserTable = "sys_message_user"

	//查询通知列表
	_queryMessageList = "select id, user_id, title, icon, content, release_time from %s where title like ? order by id desc LIMIT ?, ?"

	//查询通知数目
	_queryMessageCount = "select count(*) from %s where title like ?"

	//插入通知
	_insertMessage = "insert into %s (user_id, title, icon, content, release_time) values (?,?,?,?,?)"

	//插入用户通知
	_insertMessageUser = "insert into %s (message_id, user_id, is_read) values %s"

	//更新通知
	_updateMessage = "update %s set title = ?, icon = ?, content = ?, release_time = ? where id = ?"

	//删除通知
	_deleteMessage = "delete from %s where id = ?"
	//删除通知所有用户通知
	_deleteMessageUserByMessage = "delete from %s where message_id = ?"

	//查询用户通知简要
	_queryMessageSumList = "select id, user_id, title, icon, release_time from %s where id in (%s) and release_time < ? order by id desc"

	//查询用户通知id列表
	_queryMessageUserList = "select message_id, is_read from %s where user_id = ?"

	//查询用户通知
	_queryMessageInfo = "select id, user_id, title, icon, content, release_time from %s where id = ? and release_time < ?"

	//更新用户通知已读
	_updateMessageUserRead = "update %s set is_read = ? where message_id = ? and user_id = ?"

	//删除用户通知
	_deleteMessageUser = "delete from %s where message_id = ? and user_id = ?"

	//删除用户所有通知
	_deleteMessageUserByUser = "delete from %s where user_id = ?"
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

//查询通知分页
func (d *dao) RawMessagePage(ctx context.Context, req *pb.GetMessagePageReq) (resp *pb.GetMessagePageResp, err error) {
	sqlMessageList := fmt.Sprintf(_queryMessageList, _sysMessageTable)
	sqlMessageCount := fmt.Sprintf(_queryMessageCount, _sysMessageTable)
	query := fmt.Sprintf("%%%s%%", req.Query)
	resp = &pb.GetMessagePageResp{}
	resp.PageNum = req.PageNum
	resp.PageSize = req.PageSize
	if err = d.db.QueryRow(ctx, sqlMessageCount, query).Scan(&resp.Total); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_message]] scan all record error(%v)", err)
		return
	}
	start := (req.PageNum - 1) * req.PageSize
	end := start + req.PageSize
	rows, err := d.db.Query(ctx, sqlMessageList, query, start, end)
	if err != nil {
		log.Error("select message error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var releaseTime xtime.Time
		var userId int64
		messageInfo := &pb.MessageInfo{}
		if err = rows.Scan(&messageInfo.Id, &userId, &messageInfo.Title,
			&messageInfo.Icon, &messageInfo.Content, &releaseTime); err != nil {
			log.Error("[dao.dao-anchor.mysql|sys_message] scan short id record error(%v)", err)
			return
		}
		userNameReq := &pb.GetUserNameReq{}
		userNameReq.Id = userId
		var userNameResp *pb.GetUserNameResp
		if userNameResp, err = d.accountClient.GetUserName(ctx, userNameReq); err != nil {
			log.Error("d.accountClient.GetUserName err(%v)", err)
			return
		}
		messageInfo.Publisher = userNameResp.Name
		messageInfo.ReleaseTime = releaseTime.Time().Unix()
		resp.MessageList = append(resp.MessageList, messageInfo)
	}
	return
}

//发布通知
func (d *dao) InsertMessage(ctx context.Context, req *pb.AddMessageReq) (resp *pb.AddMessageResp, err error) {
	sqlMessage := fmt.Sprintf(_insertMessage, _sysMessageTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlMessage, req.UserId, req.Message.Title, req.Message.Icon,
		req.Message.Content, xtime.Time(req.Message.ReleaseTime).Time()); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_message]] failed to insert error(%v)", err)
		return
	}
	var messageId int64
	resp = &pb.AddMessageResp{}
	if messageId, err = res.LastInsertId(); err == nil {
		resp.Result = "success"
	}
	var data string
	for _, v := range req.UserList {
		data = data + fmt.Sprintf("(%d,%d,%d),", messageId, v, _messageUnRead)
	}
	data = data[:len(data)-1]
	sqlMessageUser := fmt.Sprintf(_insertMessageUser, _sysMessageUserTable, data)
	if res, err = d.db.Exec(ctx, sqlMessageUser); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_message_user]] failed to insert error(%v)", err)
		return
	}
	if _, err = res.LastInsertId(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//更新通知
func (d *dao) UpdateMessage(ctx context.Context, req *pb.UpdateMessageReq) (resp *pb.UpdateMessageResp, err error) {
	sqlMessage := fmt.Sprintf(_updateMessage, _sysMessageTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlMessage, req.Message.Title, req.Message.Icon, req.Message.Content,
		xtime.Time(req.Message.ReleaseTime).Time(), req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_message]] failed to update: (%v), error(%v)", req.Id, err)
	}
	resp = &pb.UpdateMessageResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//删除通知
func (d *dao) DeleteMessage(ctx context.Context, req *pb.DeleteMessageReq) (resp *pb.DeleteMessageResp, err error) {
	sqlMessage := fmt.Sprintf(_deleteMessage, _sysMessageTable)
	sqlMessageUser := fmt.Sprintf(_deleteMessageUserByMessage, _sysMessageUserTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlMessage, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_message]] failed to delete: (%v), error(%v)", req.Id, err)
		return
	}
	if res, err = d.db.Exec(ctx, sqlMessageUser, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_message_user]] failed to delete: (%v), error(%v)", req.Id, err)
		return
	}
	resp = &pb.DeleteMessageResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//查询通知简要列表
func (d *dao) RawMessageList(ctx context.Context, req *pb.GetMessageListReq) (resp *pb.GetMessageListResp, err error) {
	sqlMessageUserList := fmt.Sprintf(_queryMessageUserList, _sysMessageUserTable)
	resp = &pb.GetMessageListResp{}
	rows, err := d.db.Query(ctx, sqlMessageUserList, req.UserId)
	if err != nil {
		log.Error("select message_user error(%v)", err)
		return
	}
	defer rows.Close()
	var messages string
	readMap := make(map[int64]int32)
	for rows.Next() {
		var messageId int64
		var isRead int32
		if err = rows.Scan(&messageId, &isRead); err != nil {
			log.Error("[dao.dao-anchor.mysql|sys_message_user] scan short id record error(%v)", err)
			return
		}
		messages = messages + fmt.Sprintf("%d,", messageId)
		readMap[messageId] = isRead
	}
	if len(messages) == 0 {
		return
	}
	messages = messages[:len(messages)-1]
	sqlMessageList := fmt.Sprintf(_queryMessageSumList, _sysMessageTable, messages)
	rows, err = d.db.Query(ctx, sqlMessageList, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Error("select message error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var releaseTime xtime.Time
		var userId int64
		messageSum := &pb.MessageSum{}
		if err = rows.Scan(&messageSum.Id, &userId, &messageSum.Title,
			&messageSum.Icon, &releaseTime); err != nil {
			log.Error("[dao.dao-anchor.mysql|sys_message] scan short id record error(%v)", err)
			return
		}
		userNameReq := &pb.GetUserNameReq{}
		userNameReq.Id = userId
		var userNameResp *pb.GetUserNameResp
		if userNameResp, err = d.accountClient.GetUserName(ctx, userNameReq); err != nil {
			log.Error("d.accountClient.GetUserName err(%v)", err)
			return
		}
		messageSum.Publisher = userNameResp.Name
		messageSum.ReleaseTime = releaseTime.Time().Unix()
		messageSum.IsRead = readMap[messageSum.Id]
		resp.MessageList = append(resp.MessageList, messageSum)
	}
	return
}

//查询通知
func (d *dao) RawMessage(ctx context.Context, req *pb.GetMessageReq) (resp *pb.GetMessageResp, err error) {
	sqlMessage := fmt.Sprintf(_queryMessageInfo, _sysMessageTable)
	resp = &pb.GetMessageResp{}
	resp.MessageInfo = &pb.MessageInfo{}
	var userId int64
	var releaseTime xtime.Time
	if err = d.db.QueryRow(ctx, sqlMessage, req.Id, time.Now().Format("2006-01-02 15:04:05")).Scan(&resp.MessageInfo.Id,
		&userId, &resp.MessageInfo.Title, &resp.MessageInfo.Icon, &resp.MessageInfo.Content, &releaseTime); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_message]] scan short id record error(%v)", err)
		return
	}
	userNameReq := &pb.GetUserNameReq{}
	userNameReq.Id = userId
	var userNameResp *pb.GetUserNameResp
	if userNameResp, err = d.accountClient.GetUserName(ctx, userNameReq); err != nil {
		log.Error("d.accountClient.GetUserName err(%v)", err)
		return
	}
	resp.MessageInfo.Publisher = userNameResp.Name
	resp.MessageInfo.ReleaseTime = releaseTime.Time().Unix()
	return
}

//更新用户通知已读
func (d *dao) SetMessageUserRead(ctx context.Context, req *pb.SetMessageUserReadReq) (resp *pb.SetMessageUserReadResp, err error) {
	sqlMessage := fmt.Sprintf(_updateMessageUserRead, _sysMessageUserTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlMessage, _messageRead, req.MessageId, req.UserId); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_message_user]] failed to update: (%v), (%v),error(%v)", req.MessageId, req.UserId, err)
	}
	resp = &pb.SetMessageUserReadResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//删除用户通知
func (d *dao) DeleteMessageUser(ctx context.Context, req *pb.DeleteMessageUserReq) (resp *pb.DeleteMessageUserResp, err error) {
	sqlMessageUser := fmt.Sprintf(_deleteMessageUser, _sysMessageUserTable)
	sqlMessageUserByUser := fmt.Sprintf(_deleteMessageUserByUser, _sysMessageUserTable)
	var res dsql.Result
	if req.MessageId != 0 {
		if res, err = d.db.Exec(ctx, sqlMessageUser, req.MessageId, req.UserId); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_message]] failed to delete: (%v), error(%v)", req.UserId, err)
			return
		}
	} else {
		if res, err = d.db.Exec(ctx, sqlMessageUserByUser, req.UserId); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_message]] failed to delete: (%v), error(%v)", req.UserId, err)
			return
		}
	}
	resp = &pb.DeleteMessageUserResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}
