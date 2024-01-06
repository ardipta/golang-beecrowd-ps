/**
 * https://www.beecrowd.com.br/judge/en/problems/view/2609
 */
SELECT categories.name, sum(products.amount) from products LEFT JOIN categories on products.id_categories = categories.id GROUP by categories.name;