PRAGMA ENCODING="UTF-8";

CREATE TABLE users (
  id STRING NOT NULL,
  team_id STRING NOT NULL,
  name STRING NOT NULL,
  deleted BOOLEAN NOT NULL,
  color STRING NOT NULL,
  real_name STRING NOT NULL,
  tz STRING NOT NULL,
  tz_label STRING NOT NULL,
  tz_offset INTEGER NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE channels (
  id STRING NOT NULL,
  is_open BOOLEAN NOT NULL,
  last_read STRING NOT NULL,
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
  name_normalized STRING NOT NULL,
  num_members INTEGER NOT NULL,
--  priority DECIMAL NOT NULL,
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
  updated_at DATETIME NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE messages (
  `type` STRING NOT NULL,
  channel STRING NOT NULL,
  user STRING NOT NULL,
  `text` STRING NOT NULL,
  `timestamp` STRING NOT NULL,
  thread_timestamp STRING NOT NULL,
  is_starred BOOLEAN NOT NULL,
  last_read STRING NOT NULL,
  subscribed BOOLEAN NOT NULL,
  unread_count INTEGER NOT NULL,
  sub_type STRING NOT NULL,
  hidden BOOLEAN NOT NULL,
  deleted_timestamp STRING NOT NULL,
  event_timestamp STRING NOT NULL,
  bot_id STRING NOT NULL,
  name STRING NOT NULL,
  reply_count INTEGER NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY (`timestamp`)
);
