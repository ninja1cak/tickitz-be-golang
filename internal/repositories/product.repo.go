package repositories

import (
	"errors"
	"fmt"
	"log"
	"math"
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/models"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) CreateProduct(dataProduct *models.Product, dataSize *models.Product_size) (*config.Result, error) {

	dataProduct.Product_slug = strings.Join(strings.Split(dataProduct.Product_name, " "), "-")

	query := `INSERT INTO public.product (
		product_name, 
		product_slug,
		product_desc,
		product_stock,
		product_type,
		product_image,
		delivery_method,
		delivery_hour_start,
		delivery_hour_end		
		)
	VALUES(
		:product_name, 
		:product_slug,
		:product_desc,
		:product_stock,
		:product_type,
		:product_image,
		array[:delivery_method],
		:delivery_hour_start,
		:delivery_hour_end	
	)`

	_, err := r.NamedExec(query, dataProduct)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"product_product_slug_key\"" {
			return nil, errors.New("Duplicate Products")
		}
		return nil, err
	}

	err = r.Get(dataSize, `SELECT product_id
	FROM
		public.product
	WHERE
		product_slug = $1`, dataProduct.Product_slug)

	if err != nil {
		return nil, err
	}

	log.Println("dataSize.ProductSizeSlice", dataSize.ProductSizeSlice)
	for i := 1; i <= len(dataSize.ProductSizeSlice); i++ {
		dataSize.Product_size = strings.Join(dataSize.ProductSizeSlice[i-1:i], "")
		dataSize.Product_price = dataSize.ProductPriceSlice[i-1]
		query = `INSERT INTO public.product_size (
			product_id,
			product_size,
			product_price
			)
		VALUES(
			:product_id,
			:product_size,
			:product_price
		)`
		_, err = r.NamedExec(query, dataSize)
		if err != nil {
			return nil, err
		}
	}

	return &config.Result{Data: "product success created"}, nil
}

func (r *RepoProduct) GetProduct(limit string, page string, search string, sort string) (*config.Result, error) {
	var offset int
	var next int
	var prev int
	var qSearch string = ""
	var qSort string = ""

	if search != "" {
		search = "%" + search + "%"
		qSearch = fmt.Sprintf("AND p.product_name ILIKE '%s'", search)
	}

	if sort != "" {

		qSort = fmt.Sprintf("AND p.product_type = '%s'", sort)
	}

	counts := struct {
		Count int `db:"total"`
	}{}
	qCount := fmt.Sprintf("select count(product_id) total from product p WHERE true %s %s", qSearch, qSort)
	err := r.Get(&counts, qCount)

	if err != nil {
		return nil, err
	}
	if counts.Count == 0 {
		return &config.Result{Message: "data not found"}, nil
	}
	lim, _ := strconv.Atoi(limit)
	pag, _ := strconv.Atoi(page)
	offset = lim * (pag - 1)

	if float64(pag) == math.Ceil(float64(counts.Count)/float64(lim)) {
		next = 0
	} else {
		next = pag + 1
	}
	if pag == 1 {
		prev = 0
	} else {
		prev = pag - 1
	}

	var meta = config.Meta{
		Next:  next,
		Prev:  prev,
		Total: counts.Count,
	}

	query := fmt.Sprintf(`select 
		p.product_name, 
		p.product_desc, 
		p.product_stock,
		p.product_type,
		string_agg(ps.product_size || ' : ' || ps.product_price, ', ' order by ps.product_price, ps.product_size) size_price
	from product p
	join product_size ps 
	on p.product_id = ps.product_id %s %s
	group by p.product_id LIMIT %s OFFSET %v `, qSearch, qSort, limit, offset)

	data := []struct {
		Product_name  string `db:"product_name"`
		Product_desc  string `db:"product_desc"`
		Product_stock int    `db:"product_stock"`
		Product_type  string `db:"product_type"`
		Size_price    string `db:"size_price"`
	}{}

	err = r.Select(&data, query)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	return &config.Result{Data: data, Meta: meta}, nil
}

func (r *RepoProduct) UpdateProduct(dataProduct *models.Product, dataSize *models.Product_size) (string, error) {
	set := ""

	err := r.Get(dataProduct, `SELECT product_id
	FROM
		public.product
	WHERE
		product_slug = $1`, dataProduct.Product_slug)

	if err != nil {
		log.Println("tessssssss")
		return "", err
	}

	if dataProduct.Product_name != "" {
		set += "product_name = :product_name,"
		dataProduct.Product_slug = strings.Join(strings.Split(dataProduct.Product_name, " "), "-")
		set += "product_slug = :product_slug,"
	}

	if dataProduct.Product_desc != "" {
		set += "product_desc = :product_desc,"
	}

	if dataProduct.Product_type != "" {
		set += "product_type = :product_type,"
	}

	if dataProduct.Product_stock != 0 {
		set += "product_stock = :product_stock,"
	}

	if dataProduct.Isfavorite != false {
		set += "isfavorite = :isfavorite,"
	}

	if len(dataProduct.Delivery_method) != 0 {
		log.Println(dataProduct.Delivery_method)
		set += "delivery_method = array[:delivery_method],"
	}

	if dataProduct.Delivery_hour_start != "" {
		set += "delivery_hour_start = :delivery_hour_start,"
	}

	if dataProduct.Delivery_hour_end != "" {
		set += "delivery_hour_end = :delivery_hour_end,"
	}

	if *dataProduct.Product_image != "" {
		set += "product_image = :product_image,"
	}

	if set != "" {
		set += "updated_at = NOW()"

		query := fmt.Sprintf(`UPDATE public.product
		SET
			%s
		WHERE
			product_id = :product_id
			`, set)

		_, err = r.NamedExec(query, dataProduct)
		if err != nil {
			return "", err
		}
	}

	if len(dataSize.ProductSizeSlice) != 0 {
		dataSize.Product_id = dataProduct.Product_id
		dataSize.Product_size = strings.Join(dataSize.ProductSizeSlice[0:1], "")
		dataSize.Product_price = dataSize.ProductPriceSlice[0]
		_, err := r.NamedExec(`UPDATE public.product_size
		SET
			product_price = :product_price,
			updated_at = NOW()
		WHERE
			product_size = :product_size AND product_id = :product_id
			`, dataSize)

		if err != nil {
			return "", err
		}

	}
	return "update product success", nil
}

func (r *RepoProduct) DeleteProduct(dataProduct *models.Product) (string, error) {

	err := r.Get(dataProduct, `SELECT product_id
	FROM
		public.product
	WHERE
		product_slug = $1`, dataProduct.Product_slug)

	if err != nil {
		return "", err
	}

	_, err = r.NamedExec(`DELETE FROM public.product_size WHERE product_id = :product_id`, dataProduct)
	if err != nil {
		return "", err
	}

	_, err = r.NamedExec(`DELETE FROM public.product WHERE product_id = :product_id`, dataProduct)

	if err != nil {
		return "", err
	}

	return "delete success ", nil
}
