CREATE TABLE `group` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(255) NOT NULL COMMENT '分组名字',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `status` tinyint NOT NULL COMMENT '状态 1存在 0删除',
    PRIMARY KEY (`id`),
    UNIQUE (`name`)
);
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
);
CREATE TABLE `product` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(255) NOT NULL COMMENT '产品名',
    `characteristic` tinyint NOT NULL COMMENT '产品分类 1固定资产 0低值易耗 ',
    `quantity` int NOT NULL DEFAULT 0 COMMENT '库存数量',
    `used` int NOT NULL DEFAULT 0 COMMENT '领用或消耗的数量',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `status` tinyint NOT NULL COMMENT '状态 1存在 0删除',
    UNIQUE (`name`),
    PRIMARY KEY (`id`)
);
CREATE TABLE `record` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `staff_id` int UNSIGNED NOT NULL COMMENT '领用人ID',
    `product_id` int UNSIGNED NOT NULL COMMENT '领用物品ID',
    `count` int NOT NULL COMMENT '领用数量',
    `time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '领用时间',
     PRIMARY KEY (`id`)
);
CREATE TABLE `user` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_name` varchar(255) NOT NULL COMMENT '账户名',
    `password_digest` varchar(255) NOT NULL COMMENT '账户密码',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `status` tinyint NOT NULL COMMENT '状态 1存在 0删除',
    UNIQUE (`user_name`),
    PRIMARY KEY (`id`)
);