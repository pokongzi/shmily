CREATE TABLE IF NOT EXISTS `user` ( 
    `id` BIGINT NOT NULL auto_increment PRIMARY KEY, 
    `uid` CHAR(20) NOT NULL DEFAULT '用户id' COMMENT 'uid',
    `name` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '姓名',
    `birthday` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '生日',
    `birth_year` INT NOT NULL DEFAULT 0 COMMENT '出生年份',
    `phone` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '手机号',
    `icon` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '头像地址',
    `sex` int NOT NULL DEFAULT 0 COMMENT '性别 1男2女',
    `city` int NOT NULL DEFAULT 0 COMMENT '城市id',
    `signature` VARCHAR not NULL DEFAULT '' COMMENT '签名',
    `status` int NOT NULL DEFAULT 0 COMMENT '状态 0正常1注销',
    `create_time` int64 NOT NULL DEFAULT 0 COMMENT '创建日期',
    `update_time` int64 NOT NULL DEFAULT 0 COMMENT '更新日期'
) COMMENT '用户信息';

CREATE TABLE IF NOT EXISTS `region` (
    `id` BIGINT NOT NULL auto_increment PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '地区名称',
    `parent_id` BIGINT NOT NULL DEFAULT 0 COMMENT '父级ID，0表示省份',
    `level` TINYINT NOT NULL DEFAULT 1 COMMENT '层级 1:省份 2:城市',
    `sort_order` INT NOT NULL DEFAULT 0 COMMENT '排序',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1:启用 0:禁用',
    `create_time` int64 NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int64 NOT NULL DEFAULT 0 COMMENT '更新时间',
    INDEX `idx_parent_id` (`parent_id`),
    INDEX `idx_level` (`level`)
) COMMENT '地区信息表';

CREATE TABLE IF NOT EXISTS `user_photo` (
    `id` BIGINT NOT NULL auto_increment PRIMARY KEY,
    `uid` BIGINT NOT NULL COMMENT '用户ID',
    `photo_url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '图片URL',
    `sort_order` INT NOT NULL DEFAULT 0 COMMENT '排序',
    `photo_type` TINYINT NOT NULL DEFAULT 1 COMMENT '图片类型 1:普通照片 2:相册封面等',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1:有效 0:无效',
    `create_time` int64 NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int64 NOT NULL DEFAULT 0 COMMENT '更新时间',
    INDEX `idx_user_id` (`user_id`)
) COMMENT '用户图片表';

CREATE TABLE IF NOT EXISTS `user_hobby` (
    `id` BIGINT NOT NULL auto_increment PRIMARY KEY,
    `uid` BIGINT NOT NULL COMMENT '用户ID',
    `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '爱好名称',
    `hobby_type` TINYINT NOT NULL DEFAULT 0 COMMENT '爱好类型',
    `description` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '个人描述/补充说明',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1:有效 0:无效',
    `create_time` int64 NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int64 NOT NULL DEFAULT 0 COMMENT '更新时间',
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_hobby_id` (`hobby_id`)
) COMMENT '用户爱好关联表';

CREATE TABLE IF NOT EXISTS `user_like` (
    `id` BIGINT NOT NULL auto_increment PRIMARY KEY,
    `uid` BIGINT NOT NULL COMMENT '点赞用户ID',
    `liked_uid` BIGINT NOT NULL COMMENT '被点赞用户ID',
    `like_type` TINYINT NOT NULL DEFAULT 1 COMMENT '喜欢类型 1:普通喜欢 2:超级喜欢',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1:有效 0:取消',
    `create_time` int64 NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int64 NOT NULL DEFAULT 0 COMMENT '更新时间',
    INDEX `idx_uid` (`uid`),
    INDEX `idx_liked_uid` (`liked_uid`),
    UNIQUE INDEX `idx_user_like` (`uid`, `liked_uid`)
) COMMENT '用户点赞关系表';

CREATE TABLE IF NOT EXISTS `relation` (
    `id` BIGINT NOT NULL auto_increment PRIMARY KEY,
    `uid` CHAR(20) NOT NULL COMMENT '申请用户ID',
    `requested_uid` CHAR(20) NOT NULL COMMENT '被申请用户ID',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态 0:待处理 1:同意 2:拒绝 3:已解除',
    `request_time` int64 NOT NULL DEFAULT 0 COMMENT '申请时间',
    `response_time` int64 NOT NULL DEFAULT 0 COMMENT '对方操作时间',
    `start_time` int64 NOT NULL DEFAULT 0 COMMENT '关系开始时间',
    `end_time` int64 NOT NULL DEFAULT 0 COMMENT '关系结束时间',
    `request_msg` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '申请留言',
    `reject_reason` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '拒绝原因',
    `create_time` int64 NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int64 NOT NULL DEFAULT 0 COMMENT '更新时间',
    INDEX `idx_requester_uid` (`requester_uid`),
    INDEX `idx_requested_uid` (`requested_uid`),
    INDEX `idx_status` (`status`),
    UNIQUE INDEX `idx_relation` (`requester_uid`, `requested_uid`)
) COMMENT '用户关系申请表';


CREATE TABLE IF NOT EXISTS `user_oauth` (
    `id` BIGINT NOT NULL auto_increment PRIMARY KEY,
    `platform` TINYINT NOT NULL COMMENT '平台类型 1:微信 2:抖音 3:QQ等',
    `open_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '平台openid',
    `union_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '平台unionid',
    `nickname` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '平台昵称',
    `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '平台头像',
    `gender` TINYINT NOT NULL DEFAULT 0 COMMENT '性别 0:未知 1:男 2:女',
    `session_key` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '会话密钥',
    `access_token` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '访问令牌',
    `refresh_token` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '刷新令牌',
    `token_expire_time` int64 NOT NULL DEFAULT 0 COMMENT '令牌过期时间',
    `last_login_time` int64 NOT NULL DEFAULT 0 COMMENT '最后登录时间',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1:正常 0:禁用',
    `create_time` int64 NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int64 NOT NULL DEFAULT 0 COMMENT '更新时间',
    INDEX `idx_uid` (`uid`),
    UNIQUE INDEX `idx_platform_openid` (`platform`, `open_id`),
    INDEX `idx_union_id` (`union_id`)
) COMMENT '第三方平台用户信息表';

CREATE TABLE IF NOT EXISTS `user_oauth_bind` (
    `id` BIGINT NOT NULL auto_increment PRIMARY KEY,
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `oauth_id` BIGINT NOT NULL COMMENT '第三方平台用户ID',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1:正常 0:解绑',
    `create_time` int64 NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int64 NOT NULL DEFAULT 0 COMMENT '更新时间',
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_oauth_id` (`oauth_id`),
    UNIQUE INDEX `idx_user_oauth` (`user_id`, `oauth_id`)
) COMMENT '用户第三方平台绑定表';















