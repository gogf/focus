/*******  20201110 用户加入唯一索引  *************/
ALTER TABLE `gf`.`gf_user`
ADD UNIQUE INDEX `uni_user_passport`(`passport`) USING BTREE,
ADD UNIQUE INDEX `uni_user_nickname`(`nickname`) USING BTREE;

ALTER TABLE `gf`.`gf_content`
    ADD INDEX `idx_content_user_id`(`user_id`) USING BTREE;