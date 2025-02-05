// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatHistoriesColumns holds the columns for the "chat_histories" table.
	ChatHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "chat_id", Type: field.TypeInt64, Default: 0},
		{Name: "message_id", Type: field.TypeInt64, Default: 0},
		{Name: "user_id", Type: field.TypeInt64, Default: 0},
		{Name: "username", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "full_name", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "text", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "replied_to_message_id", Type: field.TypeInt64, Default: 0},
		{Name: "replied_to_user_id", Type: field.TypeInt64, Default: 0},
		{Name: "replied_to_full_name", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "replied_to_username", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "replied_to_text", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "chatted_at", Type: field.TypeInt64},
		{Name: "embedded", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// ChatHistoriesTable holds the schema information for the "chat_histories" table.
	ChatHistoriesTable = &schema.Table{
		Name:       "chat_histories",
		Columns:    ChatHistoriesColumns,
		PrimaryKey: []*schema.Column{ChatHistoriesColumns[0]},
	}
	// SlackOauthCredentialsColumns holds the columns for the "slack_oauth_credentials" table.
	SlackOauthCredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "team_id", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "access_token", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// SlackOauthCredentialsTable holds the schema information for the "slack_oauth_credentials" table.
	SlackOauthCredentialsTable = &schema.Table{
		Name:       "slack_oauth_credentials",
		Columns:    SlackOauthCredentialsColumns,
		PrimaryKey: []*schema.Column{SlackOauthCredentialsColumns[0]},
	}
	// TelegramChatFeatureFlagsColumns holds the columns for the "telegram_chat_feature_flags" table.
	TelegramChatFeatureFlagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "chat_id", Type: field.TypeInt64, Unique: true},
		{Name: "chat_type", Type: field.TypeString, Size: 2147483647},
		{Name: "feature_chat_histories_recap", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// TelegramChatFeatureFlagsTable holds the schema information for the "telegram_chat_feature_flags" table.
	TelegramChatFeatureFlagsTable = &schema.Table{
		Name:       "telegram_chat_feature_flags",
		Columns:    TelegramChatFeatureFlagsColumns,
		PrimaryKey: []*schema.Column{TelegramChatFeatureFlagsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatHistoriesTable,
		SlackOauthCredentialsTable,
		TelegramChatFeatureFlagsTable,
	}
)

func init() {
}
