SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ID (UUID)',
  `username` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码',
  `avatar_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户头像URL',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间 (软删除)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_username` (`username`) USING BTREE COMMENT '用户名唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频ID (UUID)',
  `user_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频作者ID',
  `video_url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频播放URL',
  `cover_url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频封面URL',
  `title` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频标题',
  `description` text COLLATE utf8mb4_general_ci COMMENT '视频描述',
  `visit_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '播放次数',
  `like_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `comment_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间 (软删除)',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '作者ID索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='视频表';

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
  `id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '点赞ID (UUID)',
  `user_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '点赞用户ID',
  `likeable_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '被点赞对象ID (视频ID或评论ID)',
  `likeable_type` enum('video','comment') COLLATE utf8mb4_general_ci NOT NULL COMMENT '被点赞对象类型 (video: 视频, comment: 评论)',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间 (软删除)',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '点赞用户ID索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='点赞表';

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论ID (UUID)',
  `user_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论用户ID',
  `video_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论视频ID',
  `parent_id` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '父评论ID (用于楼中楼)',
  `content` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
  `like_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `child_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '子评论数',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间 (软删除)',
  PRIMARY KEY (`id`),
  KEY `idx_video_id` (`video_id`) USING BTREE COMMENT '视频ID索引',
  KEY `idx_parent_id` (`parent_id`) USING BTREE COMMENT '父评论ID索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论表';

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
  `id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '关系ID (UUID)',
  `user_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '被关注用户ID',
  `follower_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '关注用户ID (粉丝ID)',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间 (软删除)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_follower` (`user_id`,`follower_id`) USING BTREE COMMENT '关注关系唯一索引',
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_follower_id` (`follower_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='关注关系表';


SET FOREIGN_KEY_CHECKS = 1;

