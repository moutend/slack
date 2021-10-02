PRAGMA ENCODING="UTF-8";

CREATE TABLE user_profile (
  user_id TEXT NOT NULL, -- not exists in API response

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

  PRIMARY KEY(user_id)
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
  priority REAL NOT NULL,
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

CREATE TABLE file (
  id TEXT NOT NULL,
  -- created
  -- timestamp
  name TEXT NOT NULL,
  title TEXT NOT NULL,
  mimetype TEXT NOT NULL,
  image_exif_rotation INTEGER NOT NULL,
  filetype TEXT NOT NULL,
  pretty_type TEXT NOT NULL,
  user TEXT NOT NULL,
  mode TEXT NOT NULL,
  editable BOOLEAN NOT NULL,
  is_external BOOLEAN NOT NULL,
  external_type TEXT NOT NULL,
  size INTEGER NOT NULL,
  url_private TEXT NOT NULL,
  url_private_download TEXT NOT NULL,
  original_h INTEGER NOT NULL,
  original_w INTEGER NOT NULL,
  thumb64 TEXT NOT NULL,
  thumb80 TEXT NOT NULL,
  thumb160 TEXT NOT NULL,
  thumb360 TEXT NOT NULL,
  thumb360_gif TEXT NOT NULL,
  thumb360_w INTEGER NOT NULL,
  thumb360_h INTEGER NOT NULL,
  thumb480 TEXT NOT NULL,
  thumb480_w INTEGER NOT NULL,
  thumb480_h INTEGER NOT NULL,
  thumb720 TEXT NOT NULL,
  thumb720_w INTEGER NOT NULL,
  thumb720_h INTEGER NOT NULL,
  thumb960 TEXT NOT NULL,
  thumb960_w INTEGER NOT NULL,
  thumb960_h INTEGER NOT NULL,
  thumb1024 TEXT NOT NULL,
  thumb1024_w INTEGER NOT NULL,
  thumb1024_h INTEGER NOT NULL,
  permalink TEXT NOT NULL,
  permalink_public TEXT NOT NULL,
  edit_link TEXT NOT NULL,
  preview TEXT NOT NULL,
  preview_highlight TEXT NOT NULL,
  lines INTEGER NOT NULL,
  lines_more INTEGER NOT NULL,
  is_public BOOLEAN NOT NULL,
  public_url_shared BOOLEAN NOT NULL,
  -- channels
  -- groups
  -- ims
  -- initial_comment
  comments_count INTEGER NOT NULL,
  num_stars INTEGER NOT NULL,
  is_starred BOOLEAN NOT NULL,
  -- shares

  created_at DATETIME NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE rel_message_file (
  message_timestamp TEXT NOT NULL,
  file_id TEXT NOT NULL,

  PRIMARY KEY(message_timestamp, file_id)
);
