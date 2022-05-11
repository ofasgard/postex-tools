package postex
//Contains functions for encrypting and decrypting payloads with XOR.

func Xorify(payload []byte, key []byte) []byte {
	output := make([]byte, len(payload))
	key_index := 0
	for index, value := range payload {
		bytekey := key[key_index]
		result := value ^ bytekey
		output[index] = result
		key_index++
		if (key_index == len(key)) {
			key_index = 0
		}
	}
	return output
}
