/**
 * https://www.beecrowd.com.br/judge/en/problems/view/2605
 */
SELECT products.name, providers.name from products LEFT JOIN providers on products.id_providers = providers.id WHERE id_categories=6;