create table if not exists `users`(
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '自增主键',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建者',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '更新者',
    `version` SMALLINT(5) unsigned NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
    `is_del` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0正常1删除',
    `user_id` char(32) NOT NULL COMMIT '用户ID',
    UNIQUE idx_user_id(user_id),
)ENGINE INNODB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';