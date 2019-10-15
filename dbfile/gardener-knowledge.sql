/*
 Navicat Premium Data Transfer

 Source Server         : postgres(localhost)
 Source Server Type    : PostgreSQL
 Source Server Version : 110005
 Source Host           : localhost:5432
 Source Catalog        : gardener-knowledge
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110005
 File Encoding         : 65001

 Date: 15/10/2019 11:32:08
*/


-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS "article";
CREATE TABLE "article" (
  "id" int4 NOT NULL DEFAULT nextval('article_id_seq'::regclass),
  "title" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 NOT NULL DEFAULT 2
)
;
ALTER TABLE "article" OWNER TO "postgres";
COMMENT ON COLUMN "article"."title" IS '文章标题';
COMMENT ON COLUMN "article"."status" IS '文章状态，2未发布';
COMMENT ON TABLE "article" IS '文章';

-- ----------------------------
-- Table structure for article_article_relation
-- ----------------------------
DROP TABLE IF EXISTS "article_article_relation";
CREATE TABLE "article_article_relation" (
  "id" int8 NOT NULL DEFAULT nextval('article_article_relation_id_seq'::regclass),
  "article_id" int4 NOT NULL,
  "relate_article_id" int4 NOT NULL,
  "position" int2 NOT NULL
)
;
ALTER TABLE "article_article_relation" OWNER TO "postgres";
COMMENT ON COLUMN "article_article_relation"."article_id" IS '文章ID';
COMMENT ON COLUMN "article_article_relation"."relate_article_id" IS '关联文章ID';
COMMENT ON COLUMN "article_article_relation"."position" IS '显示位置';
COMMENT ON TABLE "article_article_relation" IS '文章关联的文章信息';

-- ----------------------------
-- Table structure for article_fragment_relation
-- ----------------------------
DROP TABLE IF EXISTS "article_fragment_relation";
CREATE TABLE "article_fragment_relation" (
  "id" int8 NOT NULL DEFAULT nextval('article_fragment_relation_id_seq'::regclass),
  "fragment_id" int4 NOT NULL,
  "article_id" int4 NOT NULL,
  "position" int2
)
;
ALTER TABLE "article_fragment_relation" OWNER TO "postgres";
COMMENT ON COLUMN "article_fragment_relation"."fragment_id" IS '知识碎片ID';
COMMENT ON COLUMN "article_fragment_relation"."article_id" IS '文章ID';
COMMENT ON COLUMN "article_fragment_relation"."position" IS '排序信息 0 简介 1 文章内容从1开始';

-- ----------------------------
-- Table structure for article_statistics
-- ----------------------------
DROP TABLE IF EXISTS "article_statistics";
CREATE TABLE "article_statistics" (
  "id" int4 NOT NULL DEFAULT nextval('statistics_id_seq'::regclass),
  "article_id" int4 NOT NULL,
  "tag_count" int4 NOT NULL DEFAULT 0,
  "fragment_count" int4 NOT NULL DEFAULT 0
)
;
ALTER TABLE "article_statistics" OWNER TO "postgres";

-- ----------------------------
-- Table structure for article_tag_relation
-- ----------------------------
DROP TABLE IF EXISTS "article_tag_relation";
CREATE TABLE "article_tag_relation" (
  "id" int8 NOT NULL DEFAULT nextval('article_tag_relation_id_seq'::regclass),
  "article_id" int4 NOT NULL,
  "tag_article_id" int4 NOT NULL
)
;
ALTER TABLE "article_tag_relation" OWNER TO "postgres";
COMMENT ON COLUMN "article_tag_relation"."article_id" IS '文章ID';
COMMENT ON COLUMN "article_tag_relation"."tag_article_id" IS '标签ID';

-- ----------------------------
-- Table structure for detail_introduction
-- ----------------------------
DROP TABLE IF EXISTS "detail_introduction";
CREATE TABLE "detail_introduction" (
  "id" int4 NOT NULL DEFAULT nextval('"detail_introduction_id
id_seq"'::regclass),
  "summary" varchar(256) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "detail_introduction" OWNER TO "postgres";
COMMENT ON COLUMN "detail_introduction"."summary" IS '简介';
COMMENT ON COLUMN "detail_introduction"."content" IS '内容';
COMMENT ON TABLE "detail_introduction" IS '详细说明';

-- ----------------------------
-- Table structure for environment_label
-- ----------------------------
DROP TABLE IF EXISTS "environment_label";
CREATE TABLE "environment_label" (
  "id" int4 NOT NULL DEFAULT nextval('environment_label_id_seq'::regclass),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "version" varchar(32) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "environment_label" OWNER TO "postgres";
COMMENT ON COLUMN "environment_label"."name" IS '名字';
COMMENT ON COLUMN "environment_label"."version" IS '版本';
COMMENT ON TABLE "environment_label" IS '环境和版本标签';

-- ----------------------------
-- Table structure for fragment
-- ----------------------------
DROP TABLE IF EXISTS "fragment";
CREATE TABLE "fragment" (
  "id" int4 NOT NULL DEFAULT nextval('fragment_id_seq'::regclass),
  "title" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "fragment" OWNER TO "postgres";
COMMENT ON TABLE "fragment" IS '知识碎片';

-- ----------------------------
-- Table structure for fragment_introduction_relation
-- ----------------------------
DROP TABLE IF EXISTS "fragment_introduction_relation";
CREATE TABLE "fragment_introduction_relation" (
  "id" int8 NOT NULL DEFAULT nextval('fragment__id_seq'::regclass),
  "fragment_id" int4 NOT NULL,
  "detail_introduction_id" int4 NOT NULL
)
;
ALTER TABLE "fragment_introduction_relation" OWNER TO "postgres";

-- ----------------------------
-- Table structure for fragment_statistics
-- ----------------------------
DROP TABLE IF EXISTS "fragment_statistics";
CREATE TABLE "fragment_statistics" (
  "id" int4 NOT NULL DEFAULT nextval('_statistics_id_seq'::regclass),
  "fragment_id" int4 NOT NULL,
  "article_count" int4 NOT NULL DEFAULT 0,
  "tag_count" int4 NOT NULL DEFAULT 0
)
;
ALTER TABLE "fragment_statistics" OWNER TO "postgres";

-- ----------------------------
-- Table structure for fragment_tag_relation
-- ----------------------------
DROP TABLE IF EXISTS "fragment_tag_relation";
CREATE TABLE "fragment_tag_relation" (
  "id" int8 NOT NULL DEFAULT nextval('fragment_tag_relation_id_seq'::regclass),
  "fragment_id" int4 NOT NULL,
  "tag_fragment_id" int4 NOT NULL
)
;
ALTER TABLE "fragment_tag_relation" OWNER TO "postgres";
COMMENT ON COLUMN "fragment_tag_relation"."fragment_id" IS '碎片ID';
COMMENT ON COLUMN "fragment_tag_relation"."tag_fragment_id" IS '标签ID';
COMMENT ON TABLE "fragment_tag_relation" IS '文章碎片和标签关系';

-- ----------------------------
-- Table structure for introduction_environment_relation
-- ----------------------------
DROP TABLE IF EXISTS "introduction_environment_relation";
CREATE TABLE "introduction_environment_relation" (
  "id" int4 NOT NULL DEFAULT nextval('introduction_environment_relation_id_seq'::regclass),
  "detail_introduction_id" int4 NOT NULL,
  "environment_label_id" int4 NOT NULL
)
;
ALTER TABLE "introduction_environment_relation" OWNER TO "postgres";

-- ----------------------------
-- Table structure for question
-- ----------------------------
DROP TABLE IF EXISTS "question";
CREATE TABLE "question" (
  "id" int4 NOT NULL DEFAULT nextval('question_id_seq'::regclass),
  "summary" varchar(2048) COLLATE "pg_catalog"."default" NOT NULL,
  "detail" text COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "question" OWNER TO "postgres";

-- ----------------------------
-- Table structure for question_solution
-- ----------------------------
DROP TABLE IF EXISTS "question_solution";
CREATE TABLE "question_solution" (
  "id" int8 NOT NULL DEFAULT nextval('question_solution_id_seq'::regclass),
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "summary" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "question_solution" OWNER TO "postgres";

-- ----------------------------
-- Table structure for question_solution_relation
-- ----------------------------
DROP TABLE IF EXISTS "question_solution_relation";
CREATE TABLE "question_solution_relation" (
  "id" int8 NOT NULL DEFAULT nextval('question_solution_relation_id_seq'::regclass),
  "question_id" int4 NOT NULL,
  "solution_id" int4 NOT NULL,
  "solution_type" int2 NOT NULL,
  "position" int2
)
;
ALTER TABLE "question_solution_relation" OWNER TO "postgres";

-- ----------------------------
-- Table structure for tag_article
-- ----------------------------
DROP TABLE IF EXISTS "tag_article";
CREATE TABLE "tag_article" (
  "id" int4 NOT NULL DEFAULT nextval('tag_id_seq'::regclass),
  "name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL
)
;
ALTER TABLE "tag_article" OWNER TO "postgres";
COMMENT ON COLUMN "tag_article"."type" IS '类型：1 主话题，2 非主话题';

-- ----------------------------
-- Table structure for tag_article_fragment_relation
-- ----------------------------
DROP TABLE IF EXISTS "tag_article_fragment_relation";
CREATE TABLE "tag_article_fragment_relation" (
  "id" int4 NOT NULL DEFAULT nextval('tag_relation_id_seq'::regclass),
  "tag_article_id" int4 NOT NULL,
  "tag_fragment_id" int4 NOT NULL
)
;
ALTER TABLE "tag_article_fragment_relation" OWNER TO "postgres";
COMMENT ON COLUMN "tag_article_fragment_relation"."tag_article_id" IS '文章标签ID';
COMMENT ON COLUMN "tag_article_fragment_relation"."tag_fragment_id" IS '碎片标签ID';

-- ----------------------------
-- Table structure for tag_fragment
-- ----------------------------
DROP TABLE IF EXISTS "tag_fragment";
CREATE TABLE "tag_fragment" (
  "id" int4 NOT NULL DEFAULT nextval('tag_fragment_id_seq'::regclass),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL
)
;
ALTER TABLE "tag_fragment" OWNER TO "postgres";

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS "user_info";
CREATE TABLE "user_info" (
  "id" int4 NOT NULL DEFAULT nextval('user_info_id_seq'::regclass),
  "user_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "nick_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "mobile_phone" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 NOT NULL DEFAULT 1
)
;
ALTER TABLE "user_info" OWNER TO "postgres";

-- ----------------------------
-- Table structure for user_privilege
-- ----------------------------
DROP TABLE IF EXISTS "user_privilege";
CREATE TABLE "user_privilege" (
  "id" int4 NOT NULL DEFAULT nextval('user_privileges_id_seq'::regclass),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 NOT NULL DEFAULT 1
)
;
ALTER TABLE "user_privilege" OWNER TO "postgres";

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS "user_role";
CREATE TABLE "user_role" (
  "id" int4 NOT NULL DEFAULT nextval('role_id_seq'::regclass),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 NOT NULL DEFAULT 1
)
;
ALTER TABLE "user_role" OWNER TO "postgres";

-- ----------------------------
-- Table structure for user_role_privilege_relation
-- ----------------------------
DROP TABLE IF EXISTS "user_role_privilege_relation";
CREATE TABLE "user_role_privilege_relation" (
  "id" int4 NOT NULL DEFAULT nextval('user_privileges_relation_id_seq'::regclass),
  "role_id" int4 NOT NULL,
  "privilege_id" int4 NOT NULL
)
;
ALTER TABLE "user_role_privilege_relation" OWNER TO "postgres";

-- ----------------------------
-- Table structure for user_role_relation
-- ----------------------------
DROP TABLE IF EXISTS "user_role_relation";
CREATE TABLE "user_role_relation" (
  "id" int4 NOT NULL DEFAULT nextval('user_role_relation_id_seq'::regclass),
  "user_id" int4 NOT NULL,
  "role_id" int4 NOT NULL
)
;
ALTER TABLE "user_role_relation" OWNER TO "postgres";

-- ----------------------------
-- Primary Key structure for table article
-- ----------------------------
ALTER TABLE "article" ADD CONSTRAINT "article_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table article_article_relation
-- ----------------------------
ALTER TABLE "article_article_relation" ADD CONSTRAINT "article_article_relation_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table article_fragment_relation
-- ----------------------------
ALTER TABLE "article_fragment_relation" ADD CONSTRAINT "article_fragment_relation_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table article_statistics
-- ----------------------------
ALTER TABLE "article_statistics" ADD CONSTRAINT "statistics_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table article_tag_relation
-- ----------------------------
ALTER TABLE "article_tag_relation" ADD CONSTRAINT "article_tag_relation_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table detail_introduction
-- ----------------------------
ALTER TABLE "detail_introduction" ADD CONSTRAINT "detail_introduction_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table environment_label
-- ----------------------------
ALTER TABLE "environment_label" ADD CONSTRAINT "environment_label_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table fragment
-- ----------------------------
ALTER TABLE "fragment" ADD CONSTRAINT "fragment_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table fragment_introduction_relation
-- ----------------------------
ALTER TABLE "fragment_introduction_relation" ADD CONSTRAINT "fragment__pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table fragment_statistics
-- ----------------------------
ALTER TABLE "fragment_statistics" ADD CONSTRAINT "_statistics_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table fragment_tag_relation
-- ----------------------------
ALTER TABLE "fragment_tag_relation" ADD CONSTRAINT "fragment_tag_relation_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table introduction_environment_relation
-- ----------------------------
ALTER TABLE "introduction_environment_relation" ADD CONSTRAINT "introduction_environment_relation_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table question
-- ----------------------------
ALTER TABLE "question" ADD CONSTRAINT "question_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table question_solution
-- ----------------------------
ALTER TABLE "question_solution" ADD CONSTRAINT "question_solution_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table question_solution_relation
-- ----------------------------
ALTER TABLE "question_solution_relation" ADD CONSTRAINT "question_solution_relation_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table tag_article
-- ----------------------------
ALTER TABLE "tag_article" ADD CONSTRAINT "tag_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table tag_article_fragment_relation
-- ----------------------------
ALTER TABLE "tag_article_fragment_relation" ADD CONSTRAINT "tag_relation_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table tag_fragment
-- ----------------------------
ALTER TABLE "tag_fragment" ADD CONSTRAINT "tag_fragment_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_info
-- ----------------------------
ALTER TABLE "user_info" ADD CONSTRAINT "user_info_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_privilege
-- ----------------------------
ALTER TABLE "user_privilege" ADD CONSTRAINT "user_privileges_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_role
-- ----------------------------
ALTER TABLE "user_role" ADD CONSTRAINT "role_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_role_privilege_relation
-- ----------------------------
ALTER TABLE "user_role_privilege_relation" ADD CONSTRAINT "user_privileges_relation_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_role_relation
-- ----------------------------
ALTER TABLE "user_role_relation" ADD CONSTRAINT "user_role_relation_pkey" PRIMARY KEY ("id");
