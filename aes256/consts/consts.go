package consts

const (
	BLOCK_SIZE           = 16
	KEY_SIZE             = 32
	ROUND_KEY_WORD_SIZE  = 4
	ROUND_KEY_WORD_COUNT = KEY_SIZE / ROUND_KEY_WORD_SIZE
	ROUND_COUNT          = 14
	ROUND_KEYS_COUNT     = ROUND_COUNT + 1
)
