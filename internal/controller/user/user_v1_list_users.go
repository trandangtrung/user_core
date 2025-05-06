package user

import (
	"context"
	"net/url"
	"strconv"
	"time"

	v1 "github.com/quannv/strongbody-api/api/user/v1"
	"github.com/quannv/strongbody-api/internal/dto"
	userDto "github.com/quannv/strongbody-api/internal/dto/user"
	"github.com/quannv/strongbody-api/internal/entity"
	"github.com/quannv/strongbody-api/internal/storage/postgres"
)

func (c *ControllerV1) ListUsers(ctx context.Context, req *v1.ListUsersReq) (res *v1.ListUsersRes, err error) {
	urlValues := toUrlValues(req)
	userPaginateDto := dto.NewPaginationDto(urlValues)

	db := postgres.GetDatabaseConnection().Connection.Model(&entity.User{})
	db = dto.BuildWhere(db, *userPaginateDto, "users", []string{"user_name", "email", "mobile"})

	var users []entity.User
	var total int64

	db.Count(&total)
	err = db.Scopes(dto.Paginate(userPaginateDto.Page, userPaginateDto.Limit)).Find(&users).Error
	if err != nil {
		return nil, err
	}

	var usersRes []userDto.UserRes
	for _, user := range users {
		userRes := userDto.UserRes{
			ID:        user.ID,
			Email:     user.Email,
			UserName:  user.UserName,
			Mobile:    user.Mobile,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		usersRes = append(usersRes, userRes)
	}

	return &v1.ListUsersRes{
		Total:       int(total),
		TotalPage:   int(total) / int(userPaginateDto.Limit),
		CurrentPage: int(userPaginateDto.Page),
		Limit:       int(userPaginateDto.Limit),
		Data:        usersRes,
	}, nil
}

func toUrlValues(req *v1.ListUsersReq) url.Values {
	q := url.Values{}

	if req.Page != 0 {
		q.Set("page", strconv.Itoa(int(req.Page)))
	}
	if req.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(req.Limit)))
	}

	q.Set("keyword", req.Keyword)
	q.Set("order_by", req.OrderBy)
	q.Set("order_dir", req.OrderDir)

	if req.FromBirthDate != "" {
		formattedDate, _ := stringToTime(req.FromBirthDate)
		q.Set("from_birth_date", formattedDate.Format("2006-01-02"))
	}
	if req.ToBirthDate != "" {
		formattedDate, _ := stringToTime(req.ToBirthDate)
		q.Set("to_birth_date", formattedDate.Format("2006-01-02"))
	}

	return q
}

func stringToTime(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}

	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}

	return &parsedTime, nil
}

func intToString(i int) string {
	return strconv.Itoa(i)
}
