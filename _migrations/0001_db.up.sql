PRAGMA ENCODING="UTF-8";

CREATE TABLE user_profile (
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  real_name TEXT NOT NULL,
  real_name_normalized TEXT NOT NULL,
  display_name TEXT NOT NULL,
  display_name_normalized TEXT NOT NULL,
  email TEXT NOT NULL,
  skype TEXT NOT NULL,
  phone TEXT NOT NULL,
  image24 TEXT NOT NULL,
  image32 TEXT NOT NULL,
  image48 TEXT NOT NULL,
  image72 TEXT NOT NULL,
  image192 TEXT NOT NULL,
  image512 TEXT NOT NULL,
  image_original TEXT NOT NULL,
  title TEXT NOT NULL,
  bot_id TEXT NOT NULL,
  api_app_id TEXT NOT NULL,
  status_text TEXT NOT NULL,
  status_emoji TEXT NOT NULL,
  status_expiration INTEGER NOT NULL,
  team TEXT NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY(display_name_normalized)
);

CREATE TABLE user (
  id TEXT NOT NULL,
  team_id TEXT NOT NULL,
  name TEXT NOT NULL,
  deleted BOOLEAN NOT NULL,
  color TEXT NOT NULL,
  real_name TEXT NOT NULL,
  tz TEXT NOT NULL,
  tz_label TEXT NOT NULL,
  tz_offset INTEGER NOT NULL,
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
  presence TEXT NOT NULL,
  locale TEXT NOT NULL,
  -- updated
  -- enterprise

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE channel (
  -- Conversation fields
  id TEXT NOT NULL,
  -- created
  is_open BOOLEAN NOT NULL,
  last_read TEXT NOT NULL,
  -- latest
  unread_count INTEGER NOT NULL,
  unread_count_display INTEGER NOT NULL,
  is_group BOOLEAN NOT NULL,
  is_shared BOOLEAN NOT NULL,
  is_im BOOLEAN NOT NULL,
  is_ext_shared BOOLEAN NOT NULL,
  is_org_shared BOOLEAN NOT NULL,
  is_pending_ext_shared BOOLEAN NOT NULL,
  is_private BOOLEAN NOT NULL,
  is_mpim BOOLEAN NOT NULL,
  unlinked INTEGER NOT NULL,
  name_normalized TEXT NOT NULL,
  num_members INTEGER NOT NULL,
  -- priority DOUBLE NOT NULL,
  user TEXT NOT NULL,

  -- GroupConversation fields
  name TEXT NOT NULL,
  creator TEXT NOT NULL,
  is_archived BOOLEAN NOT NULL,
  -- members
  topic TEXT NOT NULL,
  purpose TEXT NOT NULL,

  -- Channel fields
  is_channel BOOLEAN NOT NULL,
  is_general BOOLEAN NOT NULL,
  is_member BOOLEAN NOT NULL,
  locale TEXT NOT NULL,

  created_at DATETIME NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE message (
  client_msg_id TEXT NOT NULL,
  `type` TEXT NOT NULL,
  channel TEXT NOT NULL,
  user TEXT NOT NULL,
  `text` TEXT NOT NULL,
  `timestamp` TEXT NOT NULL,
  thread_timestamp TEXT NOT NULL,
  is_starred BOOLEAN NOT NULL,
  -- pinned_to
  -- attachments
  -- edited
  last_read TEXT NOT NULL,
  subscribed BOOLEAN NOT NULL,
  unread_count INTEGER NOT NULL,
  sub_type TEXT NOT NULL,
  hidden BOOLEAN NOT NULL,
  deleted_timestamp TEXT NOT NULL,
  event_timestamp TEXT NOT NULL,
  bot_id TEXT NOT NULL,
  username TEXT NOT NULL,
  -- icons
  -- bot_profile
  inviter TEXT NOT NULL,
  topic TEXT NOT NULL,
  purpose TEXT NOT NULL,
  name TEXT NOT NULL,
  old_name TEXT NOT NULL,
  -- members
  reply_count INTEGER NOT NULL,
  -- replies
  parent_user_id TEXT NOT NULL,
  -- files
  upload BOOLEAN NOT NULL,
  -- comment
  item_type TEXT NOT NULL,
  reply_to INTEGER NOT NULL,
  team TEXT NOT NULL,
  -- reactions
  response_type TEXT NOT NULL,
  replace_original BOOLEAN NOT NULL,
  delete_original BOOLEAN NOT NULL,
  -- blocks

  created_at DATETIME NOT NULL,

  PRIMARY KEY (`timestamp`)
);
