CREATE TABLE `etraca_users` (
  `id`              BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT
  COMMENT '逻辑主键',
  `user_name`       VARCHAR(45)         NOT NULL DEFAULT ''
  COMMENT '用户名',
  `pass_word`       VARCHAR(100)        NOT NULL DEFAULT ''
  COMMENT '密码',
  `logon_name`      VARCHAR(100)        NOT NULL DEFAULT ''
  COMMENT '登录名',
  `last_logon_time` DATETIME            NOT NULL DEFAULT '1971-01-01 00:00:00'
  COMMENT '最后登录时间',
  `status`          INT                 NOT NULL DEFAULT 0
  COMMENT '状态',
  `create_time`     DATETIME            NOT NULL DEFAULT '1971-01-01 00:00:00'
  COMMENT '创建时间',
  `update_time`     DATETIME            NOT NULL DEFAULT '1971-01-01 00:00:00'
  COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_eu_logon_name` (`logon_name`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COMMENT = '用户信息表';

CREATE TABLE `etraca_logon_count` (
  `id`          BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT
  COMMENT '逻辑主键',
  `count`       BIGINT(20)          NOT NULL DEFAULT 0
  COMMENT '在线人数',
  `create_time` DATETIME            NOT NULL DEFAULT '1971-01-01 00:00:00'
  COMMENT '创建时间',
  `update_time` DATETIME            NOT NULL DEFAULT '1971-01-01 00:00:00'
  COMMENT '更新时间',
  `version`     BIGINT(20)          NOT NULL DEFAULT 0
  COMMENT '乐观锁',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COMMENT ='在线人数表';


INSERT INTO `etraca_logon_count` (create_time, update_time,version) VALUES (now(), now(),1);