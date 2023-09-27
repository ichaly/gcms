package auth

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/util"
	"gorm.io/gorm"
)

type ClientStore struct {
	db *gorm.DB
}

func NewOauthClientStore(d *gorm.DB) oauth2.ClientStore {
	c := &Client{}
	if !d.Migrator().HasTable(c.TableName()) {
		if err := d.Migrator().CreateTable(c); err != nil {
			panic(err)
		}
	}
	return &ClientStore{db: d}
}

type Client struct {
	Domain string `gorm:"type:varchar(512)"`
	Secret string `gorm:"type:varchar(512)"`
	Public bool
	base.Entity
}

func (Client) TableName() string {
	return "sys_client"
}

func (my *Client) GetID() string {
	return util.FormatLong(int64(my.ID))
}

func (my *Client) GetSecret() string {
	return my.Secret
}

func (my *Client) GetDomain() string {
	return my.Domain
}

func (my *Client) IsPublic() bool {
	return false
}

func (my *Client) GetUserID() string {
	return ""
}

// GetByID http://127.0.0.1:8080/oauth/token?grant_type=client_credentials&client_id=0&client_secret=eV4YeKI484E1nVoioE91Os6eOQip0TFs&scope=read
func (my *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	var c Client
	my.db.WithContext(ctx).Model(c).Where("id = ?", id).Take(&c)
	return &c, nil
}
