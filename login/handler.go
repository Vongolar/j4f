package jlogin

import (
	jauthority "JFFun/authority"
	"JFFun/data/Dcommand"
	"JFFun/data/Derror"
	"JFFun/data/Dlogin"
	jtable "JFFun/database/table"
	jserialization "JFFun/serialization"
	jtask "JFFun/task"
	"database/sql"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

//GetHandler 路由
func (m *MLogin) GetHandler() map[Dcommand.Command]func(*jtask.Task) {
	return map[Dcommand.Command]func(*jtask.Task){
		Dcommand.Command_authority:  m.authority,
		Dcommand.Command_guestLogin: m.loginByGuest,
	}
}

func (m *MLogin) authority(task *jtask.Task) {
	auth, ok := task.Data.(string)
	if !ok {
		task.Error(Derror.Error_badRequest)
		return
	}
	token, err := jwt.ParseWithClaims(auth, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(m.cfg.JwtSignKey), nil
	})
	if err != nil {
		task.Error(Derror.Error_badRequest)
		return
	}
	if cl, ok := token.Claims.(*jwt.StandardClaims); ok {
		if time.Now().Unix() > cl.ExpiresAt {
			task.Error(Derror.Error_authExpire)
			return
		}
		if len(cl.Id) == 0 {
			task.Error(Derror.Error_badRequest)
			return
		}
		task.Reply(Derror.Error_ok, cl.Id)
		return
	}
	task.Error(Derror.Error_badRequest)
	return
}

func (m *MLogin) loginByGuest(task *jtask.Task) {
	req := &Dlogin.GuestLoginReq{}
	if err := jserialization.UnMarshal(jserialization.DefaultMode, task.Raw, req); err != nil {
		task.Error(Derror.Error_badRequest)
		return
	}

	var playerID string

	//1.查询数据库
	row := m.getDB().QueryRow(fmt.Sprintf("select `pid` from `%s` where `id` = ?", jtable.Guest), req.Code)
	err := row.Scan(&playerID)
	if err != nil && err != sql.ErrNoRows {
		task.Error(Derror.Error_server)
		return
	}

	//2.创建新用户
	if len(playerID) == 0 {
		playerID = m.newPlayerID()

		//写入数据库
		tx, err := m.getDB().Begin()
		if err != nil {
			task.Error(Derror.Error_server)
			return
		}

		_, err = tx.Exec(fmt.Sprintf("insert into `%s` (`id`,`pid`) values (?,?)", jtable.Guest), req.Code, playerID)
		if err != nil {
			tx.Rollback()
			task.Error(Derror.Error_server)
			return
		}

		_, err = tx.Exec(fmt.Sprintf("insert into `%s` (`id`,`auth`) values (?,?)", jtable.Account), playerID, jauthority.Guest)
		if err != nil {
			tx.Rollback()
			task.Error(Derror.Error_server)
			return
		}
		tx.Commit()
	}

	mySigningKey := []byte(m.cfg.JwtSignKey)
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Id:        playerID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		task.Error(Derror.Error_server)
		return
	}
	task.Reply(Derror.Error_ok, &Dlogin.LoginResp{Token: ss})
}

func (m *MLogin) newPlayerID() string {
	return uuid.New().String()
}
