CREATE TABLE `group` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(255) NOT NULL COMMENT '分组名字',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `status` tinyint NOT NULL COMMENT '状态 1存在 0删除',
    PRIMARY KEY (`id`),
    UNIQUE (`name`)
)ENGINE = InnoDB CHARACTER SET = utf8mb4;
CREATE TABLE `staff` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(255) NOT NULL COMMENT '员工姓名',
    `number` char(10) NOT NULL COMMENT '工号',
    `email` varchar(255) NOT NULL COMMENT '邮箱',
    `group_id` int UNSIGNED NOT NULL COMMENT '所属分组',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `status` tinyint NOT NULL COMMENT '状态 1存在 0删除',
    PRIMARY KEY (`id`),
    UNIQUE (`number`)
)ENGINE = InnoDB CHARACTER SET = utf8mb4;
CREATE TABLE `product` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(255) NOT NULL COMMENT '产品名',
    `characteristic` tinyint NOT NULL COMMENT '产品分类 1固定资产 0低值易耗 ',
    `quantity` int NOT NULL DEFAULT 0 COMMENT '库存数量',
    `used` int NOT NULL DEFAULT 0 COMMENT '领用或消耗的数量',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `status` tinyint NOT NULL COMMENT '状态 1存在 0删除',
    UNIQUE (`name`, `characteristic`),
    PRIMARY KEY (`id`)
)ENGINE = InnoDB CHARACTER SET = utf8mb4;
CREATE TABLE `record` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `staff_id` int UNSIGNED NOT NULL COMMENT '员工ID',
    `product_id` int UNSIGNED NOT NULL COMMENT '物品ID',
    `count` int NOT NULL COMMENT '数量',
    `type` tinyint NOT NULL COMMENT '标识 1表示借出 0表示归还',
    `time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录时间',
     PRIMARY KEY (`id`)
)ENGINE = InnoDB CHARACTER SET = utf8mb4;
CREATE TABLE `user` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_name` varchar(255) NOT NULL COMMENT '账户名',
    `password_digest` varchar(255) NOT NULL COMMENT '账户密码',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `status` tinyint NOT NULL COMMENT '状态 1存在 0删除',
    UNIQUE (`user_name`),
    PRIMARY KEY (`id`)
)ENGINE = InnoDB CHARACTER SET = utf8mb4;

CREATE TABLE `fact_record` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `staff_id` int UNSIGNED NOT NULL COMMENT '员工ID',
    `product_id` int UNSIGNED NOT NULL COMMENT '物品ID',
    `count` int NOT NULL COMMENT '实际借用数量',
    `time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`),
     UNIQUE (`staff_id`, `product_id`)
)ENGINE = InnoDB CHARACTER SET = utf8mb4;