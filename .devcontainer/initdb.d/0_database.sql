CREATE DATABASE IF NOT EXISTS todo_app;
USE todo_app;

-- CREATE USER admin;
-- SET PASSWORD FOR admin = 'root';
-- GRANT SELECT, INSERT, UPDATE, DELETE, SHOW VIEW ON todo_app.* TO admin;

CREATE TABLE IF NOT EXISTS `tasks` (
  `task_id` INT NOT NULL AUTO_INCREMENT comment 'タスクID'
  , `task_name` VARCHAR(64) NOT NULL comment 'タスク名'
  -- , `task_category` TYNYINT NOT NULL comment 'タスクカテゴリ'
  , `task_deadline` DATE NOT NULL comment 'タスク〆切'
  , `task_details` VARCHAR(255) comment 'タスク詳細'
  , `created_at` DATETIME default CURRENT_TIMESTAMP NOT NULL comment '作成日時'
  , `updated_at` DATETIME default CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP comment '更新日時'
  , PRIMARY KEY (`task_id`)
  -- , CONSTRAINT `task_categories_FK1` FOREIGN KEY (`task_category`) REFERENCES `task_categories` (`id`)
) comment 'task';