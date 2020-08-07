SELECT COUNT(*)
FROM words
WHERE LENGTH(word) = $1
AND category_id = $2;
