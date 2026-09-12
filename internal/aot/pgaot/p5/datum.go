package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DatumGetExpandedArray(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v2 == int32(1) {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v5 == int32(3) {
			v15 = l0
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2))
			return v16
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, _c_F_DatumGetExpandedArray[0]))
			v11 = F_expand_array(m, l0, v9, int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = v11
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2))
				return v16
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_DatumGetExpandedArray[0]))
		v11 = F_expand_array(m, l0, v9, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2))
			return v16
		}
	}
}
func F_datum_write(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v2 = l1
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2 != 0 {
		switch l3 - int32(99) {
		case 0:
			v26 = l0
		case 1:
			v26 = (l0 + int32(7)) & int32(-8)
		default:
			v26 = (l0 + int32(1)) & int32(-2)
		case 6:
			v26 = (l0 + int32(3)) & int32(-4)
		}
		switch l4 - int32(1) {
		case 0:
			*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v2)
			v129 = v26
			v131 = int32(1)
			m.G0 = v10 + int32(16)
			return v129 + v131
		case 1:
			*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v2)
			v129 = v26
			v131 = int32(2)
			m.G0 = v10 + int32(16)
			return v129 + v131
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
				F_errmsg_internal(m, int32(_a_F_datum_write_0), v10)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_datum_write_1), int32(230), int32(_a_F_datum_write_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 3:
			*(*int32)(unsafe.Add(mBase, uint32(v26))) = v2
			v129 = v26
			v131 = int32(4)
			m.G0 = v10 + int32(16)
			return v129 + v131
		}
	} else {
		switch l4 + int32(2) {
		case 0:
			v107 = F_strlen(m, v2)
			mBase = m.M
			v109 = v107 + int32(1)
			if v109 != 0 {
				v110 = F__emscripten_memcpy_bulkmem(m, l0, v2, v109)
				mBase = m.M
			} else {
			}
			v129 = l0
			v131 = v109
			m.G0 = v10 + int32(16)
			return v129 + v131
		case 1:
			v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			if v52 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v140 = m.ExcPending
				if v140 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_datum_write_3), int32(0))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_datum_write_4), int32(2796), int32(_a_F_datum_write_5))
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if v52&int32(1) != 0 {
					v58 = int32(base.Ui32(v52) >> (uint(int32(1)) % 32))
					if v58 != 0 {
						v59 = F__emscripten_memcpy_bulkmem(m, l0, v2, v58)
						mBase = m.M
					} else {
					}
					v129 = l0
					v131 = v58
				} else {
					if l5 == int32(112) {
						switch l3 - int32(99) {
						case 0:
							v101 = l0
						case 1:
							v101 = (l0 + int32(7)) & int32(-8)
						default:
							v101 = (l0 + int32(1)) & int32(-2)
						case 6:
							v101 = (l0 + int32(3)) & int32(-4)
						}
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
						v104 = int32(base.Ui32(v102) >> (uint(int32(2)) % 32))
						if v104 != 0 {
							v105 = F__emscripten_memcpy_bulkmem(m, v101, v2, v104)
							mBase = m.M
						} else {
						}
						v129 = v101
						v131 = v104
					} else {
						if v52&int32(2) != 0 {
							switch l3 - int32(99) {
							case 0:
								v101 = l0
							case 1:
								v101 = (l0 + int32(7)) & int32(-8)
							default:
								v101 = (l0 + int32(1)) & int32(-2)
							case 6:
								v101 = (l0 + int32(3)) & int32(-4)
							}
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
							v104 = int32(base.Ui32(v102) >> (uint(int32(2)) % 32))
							if v104 != 0 {
								v105 = F__emscripten_memcpy_bulkmem(m, v101, v2, v104)
								mBase = m.M
							} else {
							}
							v129 = v101
							v131 = v104
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
							v67 = int32(base.Ui32(v65) >> (uint(int32(2)) % 32))
							v69 = v67 - int32(3)
							if base.Ui32(int32(127)) < base.Ui32(v69) {
								switch l3 - int32(99) {
								case 0:
									v101 = l0
								case 1:
									v101 = (l0 + int32(7)) & int32(-8)
								default:
									v101 = (l0 + int32(1)) & int32(-2)
								case 6:
									v101 = (l0 + int32(3)) & int32(-4)
								}
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
								v104 = int32(base.Ui32(v102) >> (uint(int32(2)) % 32))
								if v104 != 0 {
									v105 = F__emscripten_memcpy_bulkmem(m, v101, v2, v104)
									mBase = m.M
								} else {
								}
								v129 = v101
								v131 = v104
							} else {
								v72 = int32(1)
								v75 = v69<<(uint(v72)%32) | v72
								*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v75)
								v79 = int32(4)
								v82 = v67 - v79
								if v82 != 0 {
									v83 = F__emscripten_memcpy_bulkmem(m, l0+v72, v2+v79, v82)
									mBase = m.M
								} else {
								}
								v129 = l0
								v131 = v69
							}
						}
					}
				}
				m.G0 = v10 + int32(16)
				return v129 + v131
			}
		default:
			switch l3 - int32(99) {
			case 0:
				v126 = l0
			case 1:
				v126 = (l0 + int32(7)) & int32(-8)
			default:
				v126 = (l0 + int32(1)) & int32(-2)
			case 6:
				v126 = (l0 + int32(3)) & int32(-4)
			}
			if l4 != 0 {
				v127 = F__emscripten_memcpy_bulkmem(m, v126, v2, l4)
				mBase = m.M
			} else {
			}
			v129 = v126
			v131 = l4
			m.G0 = v10 + int32(16)
			return v129 + v131
		}
	}
}
