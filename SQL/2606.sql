/**
 * https://www.beecrowd.com.br/judge/en/problems/view/2606
 */
SELECT products.id, products.name from products LEFT JOIN categories on products.id_categories = categories.id where categories.name LIKE 'super%';