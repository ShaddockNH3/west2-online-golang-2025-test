CREATE TABLE `todo_list`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK, To-do list unique ID',
    `user_id`     bigint unsigned NOT NULL COMMENT 'FK, The user ID to whom this to-do list belongs',
    `title`       varchar(255) NOT NULL DEFAULT '' COMMENT 'To-do list title',
    `context`     text NULL COMMENT 'To-do list detailed content',
    `status`      tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Task status, 0: ToDo, 1: Complete',
    `created_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'To-do list create time',
    `updated_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'To-do list update time',
    `deleted_at`  timestamp NULL DEFAULT NULL COMMENT 'To-do list delete time',
    PRIMARY KEY (`id`),
    KEY           `idx_user_id` (`user_id`) COMMENT 'User ID index for quick searching'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='To-do list information table';
