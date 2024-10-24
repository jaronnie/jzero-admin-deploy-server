// Code generated by goctl. DO NOT EDIT.

package system_role_menu

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	systemRoleMenuFieldNames          = builder.RawFieldNames(&SystemRoleMenu{})
	systemRoleMenuRows                = strings.Join(systemRoleMenuFieldNames, ",")
	systemRoleMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(systemRoleMenuFieldNames, "`id`"), ",")
	systemRoleMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(systemRoleMenuFieldNames, "`id`"), "=?,") + "=?"
)

type (
	systemRoleMenuModel interface {
		Insert(ctx context.Context, data *SystemRoleMenu) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*SystemRoleMenu, error)
		Update(ctx context.Context, data *SystemRoleMenu) error
		Delete(ctx context.Context, id uint64) error

		// custom interface generated by jzero
		BulkInsert(ctx context.Context, datas []*SystemRoleMenu) error
		FindByCondition(ctx context.Context, conds ...condition.Condition) ([]*SystemRoleMenu, error)
		FindOneByCondition(ctx context.Context, conds ...condition.Condition) (*SystemRoleMenu, error)
		PageByCondition(ctx context.Context, conds ...condition.Condition) ([]*SystemRoleMenu, int64, error)
		UpdateFieldsByCondition(ctx context.Context, field map[string]any, conds ...condition.Condition) error
		DeleteByCondition(ctx context.Context, conds ...condition.Condition) error
	}

	defaultSystemRoleMenuModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SystemRoleMenu struct {
		Id         uint64        `db:"id"`
		CreateTime time.Time     `db:"create_time"`
		UpdateTime time.Time     `db:"update_time"`
		CreateBy   sql.NullInt64 `db:"create_by"`
		UpdateBy   sql.NullInt64 `db:"update_by"`
		RoleId     int64         `db:"role_id"`
		MenuId     int64         `db:"menu_id"`
	}
)

func newSystemRoleMenuModel(conn sqlx.SqlConn) *defaultSystemRoleMenuModel {
	return &defaultSystemRoleMenuModel{
		conn:  conn,
		table: "`system_role_menu`",
	}
}

func (m *defaultSystemRoleMenuModel) Delete(ctx context.Context, id uint64) error {
	sb := sqlbuilder.DeleteFrom(m.table)
	sb.Where(sb.EQ("`id`", id))
	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	return err
}

func (m *defaultSystemRoleMenuModel) FindOne(ctx context.Context, id uint64) (*SystemRoleMenu, error) {
	sb := sqlbuilder.Select(systemRoleMenuRows).From(m.table)
	sb.Where(sb.EQ("`id`", id))
	sb.Limit(1)
	sql, args := sb.Build()
	var resp SystemRoleMenu
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSystemRoleMenuModel) Insert(ctx context.Context, data *SystemRoleMenu) (sql.Result, error) {
	sql, args := sqlbuilder.NewInsertBuilder().
		InsertInto(m.table).
		Cols(systemRoleMenuRowsExpectAutoSet).
		Values(data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.RoleId, data.MenuId).Build()
	ret, err := m.conn.ExecCtx(ctx, sql, args...)
	return ret, err
}

func (m *defaultSystemRoleMenuModel) Update(ctx context.Context, data *SystemRoleMenu) error {
	sb := sqlbuilder.Update(m.table)
	split := strings.Split(systemRoleMenuRowsExpectAutoSet, ",")
	var assigns []string
	for _, s := range split {
		assigns = append(assigns, sb.Assign(s, nil))
	}
	sb.Set(assigns...)
	sb.Where(sb.EQ("`id`", nil))
	sql, _ := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.RoleId, data.MenuId, data.Id)
	return err
}

func (m *defaultSystemRoleMenuModel) tableName() string {
	return m.table
}

func (m *customSystemRoleMenuModel) BulkInsert(ctx context.Context, datas []*SystemRoleMenu) error {
	sb := sqlbuilder.InsertInto(m.table)
	sb.Cols(systemRoleMenuRowsExpectAutoSet)
	for _, data := range datas {
		sb.Values(data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.RoleId, data.MenuId)
	}
	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	return err
}

func (m *customSystemRoleMenuModel) FindByCondition(ctx context.Context, conds ...condition.Condition) ([]*SystemRoleMenu, error) {
	sb := sqlbuilder.Select(systemRoleMenuFieldNames...).From(m.table)
	condition.ApplySelect(sb, conds...)
	sql, args := sb.Build()

	var resp []*SystemRoleMenu
	err := m.conn.QueryRowsCtx(ctx, &resp, sql, args...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customSystemRoleMenuModel) FindOneByCondition(ctx context.Context, conds ...condition.Condition) (*SystemRoleMenu, error) {
	sb := sqlbuilder.Select(systemRoleMenuFieldNames...).From(m.table)
	condition.ApplySelect(sb, conds...)
	sb.Limit(1)
	sql, args := sb.Build()

	var resp SystemRoleMenu
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *customSystemRoleMenuModel) PageByCondition(ctx context.Context, conds ...condition.Condition) ([]*SystemRoleMenu, int64, error) {
	sb := sqlbuilder.Select(systemRoleMenuFieldNames...).From(m.table)
	countsb := sqlbuilder.Select("count(*)").From(m.table)

	condition.ApplySelect(sb, conds...)

	var countConds []condition.Condition
	for _, cond := range conds {
		if cond.Operator != condition.Limit && cond.Operator != condition.Offset {
			countConds = append(countConds, cond)
		}
	}
	condition.ApplySelect(countsb, countConds...)

	var resp []*SystemRoleMenu

	sql, args := sb.Build()
	err := m.conn.QueryRowsCtx(ctx, &resp, sql, args...)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	sql, args = countsb.Build()
	err = m.conn.QueryRowCtx(ctx, &total, sql, args...)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func (m *customSystemRoleMenuModel) UpdateFieldsByCondition(ctx context.Context, field map[string]any, conds ...condition.Condition) error {
	if field == nil {
		return nil
	}

	sb := sqlbuilder.Update(m.table)
	condition.ApplyUpdate(sb, conds...)

	var assigns []string
	for key, value := range field {
		assigns = append(assigns, sb.Assign(key, value))
	}
	sb.Set(assigns...)

	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (m *customSystemRoleMenuModel) DeleteByCondition(ctx context.Context, conds ...condition.Condition) error {
	if len(conds) == 0 {
		return nil
	}
	sb := sqlbuilder.DeleteFrom(m.table)
	condition.ApplyDelete(sb, conds...)
	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	return err
}
