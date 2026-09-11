package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tbm_create(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int64
	_ = v27
	v6 = F_palloc0(m, int32(116))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(478)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_tbm_create[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+112)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(0)
		v17 = int32(16)
		v19 = base.I32_div_u_s(l0, int32(56))
		if base.Ui32(v19) <= base.Ui32(v17) {
			v22 = v17
		} else {
			v22 = v19
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v13
		v27 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+96)) = v27
		*(*int64)(unsafe.Add(mBase, uint32(v6)+104)) = v27
		return v6
	}
}
func F_tbm_extract_page_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = v3
	v14 = v3
	for {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8)+v14<<(uint(int32(2))%32))))
		if v21 != 0 {
			v26 = v21
			v28 = v13
			v30 = v14<<(uint(int32(5))%32) | int32(1)
			for {
				if v26&int32(1) != 0 {
					if base.Ui32(v28) < base.Ui32(int32(291)) {
						*(*uint16)(unsafe.Add(mBase, uint32(l1+v28<<(uint(int32(1))%32)))) = uint16(v30)
					} else {
					}
					v43 = v28 + int32(1)
				} else {
					v43 = v28
				}
				v44 = int32(1)
				if base.Ui32(v44) < base.Ui32(v26) {
					v26 = int32(base.Ui32(v26) >> (uint(v44) % 32))
					v28 = v43
					v30 = v30 + v44
					continue
				} else {
					break
				}
				break
			}
			v52 = v43
		} else {
			v52 = v13
		}
		v58 = v14 + int32(1)
		if v58 != int32(10) {
			v13 = v52
			v14 = v58
			continue
		} else {
			break
		}
		break
	}
	return v52
}
func F_tbm_get_pageentry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v12 {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
		v33 = l0 + int32(40)
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
		if v33&int32(3) == int32(0) {
			v40 = v33 + int32(48)
			if base.Ui32(v40) <= base.Ui32(v33) {
			} else {
				v46 = v33 + int32(4)
				if base.Ui32(v46) < base.Ui32(v40) {
					v48 = v40
				} else {
					v48 = v46
				}
				v55 = F__emscripten_memset_bulkmem(m, v33, base.I32_extend8_s(int32(0)), (v33^int32(-1)+v48)&int32(-4)+int32(4))
				mBase = m.M
			}
		} else {
			v56 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v33)+5)) = v56
			*(*int32)(unsafe.Add(mBase, uint32(v33)+44)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v33)+37)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v33)+29)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v33)+21)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v33)+13)) = v56
		}
		*(*int32)(unsafe.Add(mBase, uint32(v33))) = l1
		*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)) = uint8(v34)
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v73 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v72 + v73
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v76 + v73
		v80 = v33
		m.G0 = v10 + int32(16)
		return v80
	case 1:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if l1 == v17 {
			v80 = l0 + int32(40)
			m.G0 = v10 + int32(16)
			return v80
		} else {
			F_tbm_create_pagetable(m, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v28 = F_pagetable_insert(m, v25, l1, v10+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v30&int32(1) != 0 {
						v80 = v28
					} else {
						v33 = v28
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
						if v33&int32(3) == int32(0) {
							v40 = v33 + int32(48)
							if base.Ui32(v40) <= base.Ui32(v33) {
							} else {
								v46 = v33 + int32(4)
								if base.Ui32(v46) < base.Ui32(v40) {
									v48 = v40
								} else {
									v48 = v46
								}
								v55 = F__emscripten_memset_bulkmem(m, v33, base.I32_extend8_s(int32(0)), (v33^int32(-1)+v48)&int32(-4)+int32(4))
								mBase = m.M
							}
						} else {
							v56 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v33)+5)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v33)+44)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v33)+37)) = v56
							*(*int64)(unsafe.Add(mBase, uint32(v33)+29)) = v56
							*(*int64)(unsafe.Add(mBase, uint32(v33)+21)) = v56
							*(*int64)(unsafe.Add(mBase, uint32(v33)+13)) = v56
						}
						*(*int32)(unsafe.Add(mBase, uint32(v33))) = l1
						*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)) = uint8(v34)
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v73 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v72 + v73
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v76 + v73
						v80 = v33
					}
					m.G0 = v10 + int32(16)
					return v80
				}
			}
		}
	default:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v28 = F_pagetable_insert(m, v25, l1, v10+int32(15))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
			if v30&int32(1) != 0 {
				v80 = v28
			} else {
				v33 = v28
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
				if v33&int32(3) == int32(0) {
					v40 = v33 + int32(48)
					if base.Ui32(v40) <= base.Ui32(v33) {
					} else {
						v46 = v33 + int32(4)
						if base.Ui32(v46) < base.Ui32(v40) {
							v48 = v40
						} else {
							v48 = v46
						}
						v55 = F__emscripten_memset_bulkmem(m, v33, base.I32_extend8_s(int32(0)), (v33^int32(-1)+v48)&int32(-4)+int32(4))
						mBase = m.M
					}
				} else {
					v56 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v33)+5)) = v56
					*(*int32)(unsafe.Add(mBase, uint32(v33)+44)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v33)+37)) = v56
					*(*int64)(unsafe.Add(mBase, uint32(v33)+29)) = v56
					*(*int64)(unsafe.Add(mBase, uint32(v33)+21)) = v56
					*(*int64)(unsafe.Add(mBase, uint32(v33)+13)) = v56
				}
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = l1
				*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)) = uint8(v34)
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v73 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v72 + v73
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v76 + v73
				v80 = v33
			}
			m.G0 = v10 + int32(16)
			return v80
		}
	}
}
