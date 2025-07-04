package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSet(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     []int
		wantValue []int
	}{
		{
			name:      "one element",
			key:       "key1",
			value:     []int{3},
			wantValue: []int{3},
		},
		{
			name:      "several elements",
			key:       "key2",
			value:     []int{1, 2, 3},
			wantValue: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeMap()
			for _, v := range tt.value {
				s.MapSet(tt.key, v)
			}
			result, ok := s.MapGet(tt.key)
			assert.True(t, ok, "Key not found")
			assert.Equal(t, tt.wantValue, result)
		})
	}
}

func TestMapGet(t *testing.T) {
	tests := []struct {
		name       string
		setupKey   string
		getKey     string
		values     []int
		wantValue  []int
		wantExists bool
	}{
		{
			name:       "existing key",
			setupKey:   "key 1",
			getKey:     "key 1",
			values:     []int{1, 2, 3},
			wantValue:  []int{1, 2, 3},
			wantExists: true,
		},
		{
			name:       "not existing key",
			setupKey:   "key 1",
			getKey:     "key 2",
			values:     []int{1, 2, 3},
			wantValue:  nil,
			wantExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeMap()
			for _, v := range tt.values {
				s.MapSet(tt.setupKey, v)
			}

			result, ok := s.MapGet(tt.getKey)
			assert.Equal(t, tt.wantExists, ok)
			assert.Equal(t, tt.wantValue, result)
		})
	}
}

func TestMapDelete(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		setupKey    string
		getKey      string
		values      []int
		wantExists  bool
		wantDeleted bool
	}{
		{
			name:        "valid delition",
			setupKey:    "key 1",
			getKey:      "key 1",
			values:      []int{1, 2, 3},
			wantExists:  true,
			wantDeleted: true,
		},
		{
			name:        "key not found",
			setupKey:    "key 1",
			getKey:      "key 2",
			values:      []int{1, 2, 3},
			wantExists:  false,
			wantDeleted: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeMap()
			//заполняем значениями
			for _, v := range tt.values {
				s.MapSet(tt.setupKey, v)
			}
			_, ok := s.MapGet(tt.getKey)       //проверка существует ли ключ
			assert.Equal(t, tt.wantExists, ok) //ошибка ключ не найден
			//если существует
			if ok {
				s.MapDelete(tt.getKey) //удаление улюча

				_, ok = s.MapGet(tt.getKey)          //проверка что удалили
				assert.Equal(t, tt.wantDeleted, !ok) //сравнение. !ок потому что MapGet возвращает false, а нужно true - удалено
			}
		})
	}
}
