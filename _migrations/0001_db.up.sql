PRAGMA ENCODING="UTF-8";

CREATE TABLE user_profile (
  first_name STRING NOT NULL,
  last_name STRING NOT NULL,
  real_name STRING NOT NULL,
  real_name_normalized STRING NOT NULL,
  display_name STRING NOT NULL,
  display_name_normalized STRING NOT NULL,
  email STRING NOT NULL,
  skype STRING NOT NULL,
  phone STRING NOT NULL,
  image24 STRING NOT NULL,
  image32 STRING NOT NULL,
  image48 STRING NOT NULL,
  image72 STRING NOT NULL,
  image192 STRING NOT NULL,
  image512 STRING NOT NULL,
  image_original STRING NOT NULL,
  title STRING NOT NULL,
  bot_id STRING NOT NULL,
  api_app_id STRING NOT NULL,
  status_text STRING NOT NULL,
  status_emoji STRING NOT NULL,
  status_expiration INT NOT NULL,
  team STRING NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY(display_name_normalized)
);

CREATE TABLE user (
  id STRING NOT NULL,
  team_id STRING NOT NULL,
  name STRING NOT NULL,
  deleted BOOLEAN NOT NULL,
  color STRING NOT NULL,
  real_name STRING NOT NULL,
  tz STRING NOT NULL,
  tz_label STRING NOT NULL,
  tz_offset INT NOT NULL,
  -- profile
  is_bot BOOLEAN NOT NULL,
  is_admin BOOLEAN NOT NULL,
  is_owner BOOLEAN NOT NULL,
  is_primary_owner BOOLEAN NOT NULL,
  is_restricted BOOLEAN NOT NULL,
  is_ultra_restricted BOOLEAN NOT NULL,
  is_stranger BOOLEAN NOT NULL,
  is_app_user BOOLEAN NOT NULL,
  is_invited_user BOOLEAN NOT NULL,
  has_2fa BOOLEAN NOT NULL,
  has_files BOOLEAN NOT NULL,
  presence STRING NOT NULL,
  locale STRING NOT NULL,
  -- updated
  -- enterprise

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE channel (
  id STRING NOT NULL,
  is_open BOOLEAN NOT NULL,
  last_read STRING NOT NULL,
  unread_count INT NOT NULL,
  unread_count_display INT NOT NULL,
  is_group BOOLEAN NOT NULL,
  is_shared BOOLEAN NOT NULL,
  is_im BOOLEAN NOT NULL,
  is_ext_shared BOOLEAN NOT NULL,
  is_org_shared BOOLEAN NOT NULL,
  is_pending_ext_shared BOOLEAN NOT NULL,
  is_private BOOLEAN NOT NULL,
  is_mpim BOOLEAN NOT NULL,
  unlinked INT NOT NULL,
  name_normalized STRING NOT NULL,
  num_members INT NOT NULL,
  user STRING NOT NULL,
  name STRING NOT NULL,
  creator STRING NOT NULL,
  is_archived BOOLEAN NOT NULL,
  topic STRING NOT NULL,
  purpose STRING NOT NULL,
  is_channel BOOLEAN NOT NULL,
  is_general BOOLEAN NOT NULL,
  is_member BOOLEAN NOT NULL,
  locale STRING NOT NULL,

  created_at DATETIME NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE message (
  client_msg_id STRING NOT NULL,
  `type` STRING NOT NULL,
  channel STRING NOT NULL,
  user STRING NOT NULL,
  `text` STRING NOT NULL,
  `timestamp` STRING NOT NULL,
  thread_timestamp STRING NOT NULL,
  is_starred BOOLEAN NOT NULL,
  -- pinned_to
  -- attachments
  -- edited
  last_read STRING NOT NULL,
  subscribed BOOLEAN NOT NULL,
  unread_count INT NOT NULL,
  sub_type STRING NOT NULL,
  hidden BOOLEAN NOT NULL,
  deleted_timestamp STRING NOT NULL,
  event_timestamp STRING NOT NULL,
  bot_id STRING NOT NULL,
  user_name STRING NOT NULL,
  -- icons
  -- bot_profile
  inviter STRING NOT NULL,
  topic STRING NOT NULL,
  purpose STRING NOT NULL,
  name STRING NOT NULL,
  old_name STRING NOT NULL,
  -- members
  reply_count INT NOT NULL,
  -- replies
  parent_user_id STRING NOT NULL,
  -- files
  upload BOOLEAN NOT NULL,
  -- comment
  item_type STRING NOT NULL,
  reply_to INT NOT NULL,
  team STRING NOT NULL,
  -- reactions
  response_type STRING NOT NULL,
  replace_original BOOLEAN NOT NULL,
  delete_original BOOLEAN NOT NULL,
  -- blocks

  created_at DATETIME NOT NULL,

  PRIMARY KEY (`timestamp`)
);
