SELECT word
FROM words
WHERE LENGTH(word) = $1
AND category_id = $2
OFFSET $3
LIMIT 1;
