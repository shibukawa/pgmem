package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlabFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v2 = int32(0)
	v8 = l0 - int32(8)
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v15 = v8 - base.I32_wrap_i64(int64(base.Ui64(v9)>>(uint(int64(34))%64)))&int32(1073741822)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v8
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v21 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v20 + v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v27 = v20>>(uint(v24)%32) + v21
	v32 = v2 - (v2-v20)>>(uint(v24)%32)
	if v27 == v32 {
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v35
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v35))) = v37
		v41 = v16 + v27<<(uint(int32(3))%32)
		v43 = v41 + int32(80)
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+84))
		if v44 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v43))) = v43
			v48 = v43
		} else {
			v48 = v44
		}
		*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v43
		*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v48
		v52 = v15 + int32(20)
		*(*int32)(unsafe.Add(mBase, uint32(v48))) = v52
		*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v52
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
		if v55 < v32 {
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
			if v57 != 0 {
				if v57 != v16+int32(88) {
					v71 = int32(1)
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
					if v62 != 0 {
						v69 = base.B2i32(v62 != v16+int32(96)) << (uint(int32(1)) % 32)
					} else {
						v69 = int32(0)
					}
					v71 = v69
				}
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
				if v62 != 0 {
					v69 = base.B2i32(v62 != v16+int32(96)) << (uint(int32(1)) % 32)
				} else {
					v69 = int32(0)
				}
				v71 = v69
			}
			*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v71
		}
	}
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v16)+56))
	if v76 != v77 {
	} else {
		v79 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
		v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v80
		v82 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v80))) = v82
		v84 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
		if base.Ui32(v84) <= base.Ui32(int32(9)) {
			v88 = v16 + int32(68)
			v90 = v15 + int32(20)
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
			if v91 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v16 + int32(68)
				v99 = v88
			} else {
				v99 = v91
			}
			*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v88
			*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v99
			*(*int32)(unsafe.Add(mBase, uint32(v99))) = v90
			*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v90
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v104 + int32(1)
		} else {
			F_emscripten_builtin_free(m, v15)
			mBase = m.M
			v109 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			v110 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v109 - v110
		}
		v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
		if v116 != v27 {
		} else {
			v122 = v16 + v27<<(uint(int32(3))%32) + int32(80)
			v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
			if v123 != v122 {
				v126 = v123
			} else {
				v126 = int32(0)
			}
			if v126 != 0 {
			} else {
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
				if v127 != 0 {
					if v127 != v16+int32(88) {
						v141 = int32(1)
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
						if v132 != 0 {
							v139 = base.B2i32(v132 != v16+int32(96)) << (uint(int32(1)) % 32)
						} else {
							v139 = int32(0)
						}
						v141 = v139
					}
				} else {
					v132 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
					if v132 != 0 {
						v139 = base.B2i32(v132 != v16+int32(96)) << (uint(int32(1)) % 32)
					} else {
						v139 = int32(0)
					}
					v141 = v139
				}
				*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v141
			}
		}
	}
	return
}
