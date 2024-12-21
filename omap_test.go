package omap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOmap(t *testing.T) {
	t.Run("main", func(t *testing.T) {
		om := New[string, int]().WithCap(10)

		// Set
		om.Set("test", 1)

		// Get
		val, ok := om.Get("test")
		assert.True(t, ok)
		assert.Equal(t, val, 1)

		// Get zero
		val, ok = om.Get("test1")
		assert.False(t, ok)
		assert.Equal(t, val, 0)

		// Len
		assert.Equal(t, 1, om.Len())

		// Delete
		ok = om.Delete("test")
		assert.True(t, ok)
		assert.Equal(t, 0, om.Len())

		// No delete
		ok = om.Delete("test")
		assert.False(t, ok)
		assert.Equal(t, 0, om.Len())
	})

	t.Run("iter", func(t *testing.T) {
		om := New[string, int]()

		// Cases
		cases := map[int]struct {
			expectedKey   string
			expectedValue int
		}{
			0: {
				expectedKey:   "test",
				expectedValue: 1,
			},
			1: {
				expectedKey:   "test2",
				expectedValue: 10,
			},
			2: {
				expectedKey:   "test3",
				expectedValue: 20,
			},
			3: {
				expectedKey:   "test77",
				expectedValue: 1000,
			},
		}

		// Set
		for i := 0; i < len(cases); i++ {
			om.Set(cases[i].expectedKey, cases[i].expectedValue)
		}
		assert.Equal(t, 4, om.Len())

		t.Run("forward", func(t *testing.T) {
			cnt := 0
			om.Iter(func(key string, value int) bool {
				assert.Equal(t, cases[cnt].expectedKey, key)
				assert.Equal(t, cases[cnt].expectedValue, value)
				cnt++
				return true
			})
			assert.Equal(t, 4, cnt)
		})

		t.Run("backward", func(t *testing.T) {
			cnt := 3
			om.IterBack(func(key string, value int) bool {
				assert.Equal(t, cases[cnt].expectedKey, key)
				assert.Equal(t, cases[cnt].expectedValue, value)
				cnt--
				return true
			})
			assert.Equal(t, -1, cnt)
		})
	})

	t.Run("reset middle", func(t *testing.T) {
		om := New[string, int]().WithCap(10)

		// Set
		om.Set("test", 1)
		om.Set("test2", 10)
		om.Set("test3", 12)

		// Len
		assert.Equal(t, 3, om.Len())

		// Reset
		om.Set("test2", 100)

		// Get
		v, ok := om.Get("test2")
		assert.True(t, ok)
		assert.Equal(t, 100, v)

		// Len
		assert.Equal(t, 3, om.Len())

		// Cases
		cases := map[int]struct {
			expectedKey   string
			expectedValue int
		}{
			0: {
				expectedKey:   "test",
				expectedValue: 1,
			},
			1: {
				expectedKey:   "test3",
				expectedValue: 12,
			},
			2: {
				expectedKey:   "test2",
				expectedValue: 100,
			},
		}

		t.Run("forward", func(t *testing.T) {
			cnt := 0
			om.Iter(func(key string, value int) bool {
				assert.Equal(t, cases[cnt].expectedKey, key)
				assert.Equal(t, cases[cnt].expectedValue, value)
				cnt++
				return true
			})
			assert.Equal(t, 3, cnt)
		})

		t.Run("backward", func(t *testing.T) {
			cnt := 2
			om.IterBack(func(key string, value int) bool {
				assert.Equal(t, cases[cnt].expectedKey, key)
				assert.Equal(t, cases[cnt].expectedValue, value)
				cnt--
				return true
			})
			assert.Equal(t, -1, cnt)
		})
	})

	t.Run("replace middle", func(t *testing.T) {
		om := New[string, int]().WithCap(10)

		// Set
		om.Set("test", 1)
		om.Set("test2", 10)
		om.Set("test3", 12)

		// Len
		assert.Equal(t, 3, om.Len())

		// Replace
		ok := om.Replace("test2", 100)
		assert.True(t, ok)

		// Get
		v, ok := om.Get("test2")
		assert.True(t, ok)
		assert.Equal(t, 100, v)

		// Len
		assert.Equal(t, 3, om.Len())

		// Cases
		cases := map[int]struct {
			expectedKey   string
			expectedValue int
		}{
			0: {
				expectedKey:   "test",
				expectedValue: 1,
			},
			1: {
				expectedKey:   "test2",
				expectedValue: 100,
			},
			2: {
				expectedKey:   "test3",
				expectedValue: 12,
			},
		}

		t.Run("forward", func(t *testing.T) {
			cnt := 0
			om.Iter(func(key string, value int) bool {
				assert.Equal(t, cases[cnt].expectedKey, key)
				assert.Equal(t, cases[cnt].expectedValue, value)
				cnt++
				return true
			})
			assert.Equal(t, 3, cnt)
		})

		t.Run("backward", func(t *testing.T) {
			cnt := 2
			om.IterBack(func(key string, value int) bool {
				assert.Equal(t, cases[cnt].expectedKey, key)
				assert.Equal(t, cases[cnt].expectedValue, value)
				cnt--
				return true
			})
			assert.Equal(t, -1, cnt)
		})
	})

	t.Run("delete corner", func(t *testing.T) {
		om := New[string, int]().WithCap(10)

		// Set
		om.Set("test", 1)
		om.Set("test2", 10)
		om.Set("test3", 12)
		om.Set("test4", 123)

		// Len
		assert.Equal(t, 4, om.Len())

		// Delete
		ok := om.Delete("test")
		assert.True(t, ok)

		ok = om.Delete("test4")
		assert.True(t, ok)

		// Len
		assert.Equal(t, 2, om.Len())

		// Cases
		cases := map[int]struct {
			expectedKey   string
			expectedValue int
		}{
			0: {
				expectedKey:   "test2",
				expectedValue: 10,
			},
			1: {
				expectedKey:   "test3",
				expectedValue: 12,
			},
		}

		t.Run("forward", func(t *testing.T) {
			cnt := 0
			om.Iter(func(key string, value int) bool {
				assert.Equal(t, cases[cnt].expectedKey, key)
				assert.Equal(t, cases[cnt].expectedValue, value)
				cnt++
				return true
			})
			assert.Equal(t, 2, cnt)
		})

		t.Run("backward", func(t *testing.T) {
			cnt := 1
			om.IterBack(func(key string, value int) bool {
				assert.Equal(t, cases[cnt].expectedKey, key)
				assert.Equal(t, cases[cnt].expectedValue, value)
				cnt--
				return true
			})
			assert.Equal(t, -1, cnt)
		})
	})

	t.Run("errors", func(t *testing.T) {
		var om *Omap[string, string]

		// Get
		v, ok := om.Get("test")
		assert.False(t, ok)
		assert.Equal(t, "", v)

		// Delete
		ok = om.Delete("test")
		assert.False(t, ok)

		// Len
		assert.Equal(t, 0, om.Len())

		t.Run("panics", func(t *testing.T) {
			t.Run("set", func(t *testing.T) {
				defer func() {
					e := recover()
					err, ok := e.(error)
					assert.True(t, ok)
					assert.ErrorIs(t, ErrIsntInited, err)
				}()

				om.Set("test", "test")
			})
			t.Run("replace", func(t *testing.T) {
				defer func() {
					e := recover()
					err, ok := e.(error)
					assert.True(t, ok)
					assert.ErrorIs(t, ErrIsntInited, err)
				}()

				om.Replace("test", "test")
			})
			t.Run("iter", func(t *testing.T) {
				defer func() {
					e := recover()
					err, ok := e.(error)
					assert.True(t, ok)
					assert.ErrorIs(t, ErrIsntInited, err)
				}()

				om.Iter(func(key string, value string) bool {
					return true
				})
			})

			t.Run("iter back", func(t *testing.T) {
				defer func() {
					e := recover()
					err, ok := e.(error)
					assert.True(t, ok)
					assert.ErrorIs(t, ErrIsntInited, err)
				}()

				om.IterBack(func(key string, value string) bool {
					return true
				})
			})
		})
	})

}
