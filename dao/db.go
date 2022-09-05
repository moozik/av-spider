package dao

import (
	"av-spider/entity"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

var (
	gormInstance *gorm.DB
	dbOnce       sync.Once
)

func GetDb() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		gormInstance, err = gorm.Open(sqlite.Open(`E:\github\av-spider\avmoo_test_new.db`), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return gormInstance
}

func SearchWhereBuild(keyword string) (where []string) {
	where = []string{}
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return where
	}
	for _, item := range strings.Split(keyword, " ") {
		if item == "" {
			continue
		}
		if item == "已发布" {
			where = append(where, fmt.Sprintf(`av_list.release_date <= '%s'`, time.Now().Format(`2006-01-02`)))
			continue
		}
		if item == "有资源" {
			where = append(where, `av_id IN (SELECT distinct key FROM av_extend WHERE class='movie_res')`)
			continue
		}
		if item == "已下载" {
			where = append(where, `av_id IN (SELECT distinct key FROM av_extend WHERE class='movie_res' AND val LIKE '_:\\%')`)
			continue
		}
		where = append(where, SearchWhere(item))
	}
	return
}

func SearchWhere(keyword string) string {
	return fmt.Sprintf(`(av_list.title LIKE '%%%s%%' OR 
	av_list.genre LIKE '%%%s%%' OR 
	av_list.actress LIKE '%%%s%%')`, keyword, keyword, keyword)
}

func SelectAvList(whereList []string, pageNum int) (ret []entity.AvList) {
	db := GetDb()
	db.Table("av_list").Order("release_date DESC,av_id DESC").Limit(30).Offset(pageNum)
	for _, whereItem := range whereList {
		db.Where(whereItem)
	}
	db.Scan(&ret)
	return
}

/*
# 查询列表
def select_av_list(av_list_where: list, page_num: int):
    # 每页展示的数量
    page_limit = CONFIG.get("website.page_limit")

    sql_order_by = "release_date DESC,av_id DESC"
    sql_text = ""
    if non_empty(av_list_where):
        where_str = " AND ".join(av_list_where)
        sql_text = "SELECT * FROM av_list WHERE {}".format(where_str)
    else:
        sql_text = "SELECT * FROM av_list"

    result = query_sql(
        sql_text + ' ORDER BY {} LIMIT {},{}'.format(sql_order_by, (page_num - 1) * page_limit, page_limit))

    for i in range(len(result)):
        # 扩展信息
        extend_list = storage(AV_EXTEND, {"class": "movie_res", "key": [result[i]['av_id']]}, "val")
        if not extend_list:
            continue
        for extend in extend_list:
            extend = url_rename(extend)
            if extend[:6] == "magnet" or extend[:3] == "115":
                result[i]['magnet'] = 1
                continue
            if extend[:4] == "http":
                result[i]['http'] = 1
                continue
            if extend[1] == ":":
                result[i]['file'] = 1
    return result, get_sql_count(sql_text)
*/
