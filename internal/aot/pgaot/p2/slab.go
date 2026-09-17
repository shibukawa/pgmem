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
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
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
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
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
					v72 = int32(1)
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
					if v64 != v16+int32(96) {
						v68 = int32(2)
					} else {
						v68 = int32(0)
					}
					if v64 != 0 {
						v70 = v68
					} else {
						v70 = int32(0)
					}
					v72 = v70
				}
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
				if v64 != v16+int32(96) {
					v68 = int32(2)
				} else {
					v68 = int32(0)
				}
				if v64 != 0 {
					v70 = v68
				} else {
					v70 = int32(0)
				}
				v72 = v70
			}
			*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v72
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
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
			if v89 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v88
				v95 = v88
			} else {
				v95 = v89
			}
			*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v88
			*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v95
			v99 = v15 + int32(20)
			*(*int32)(unsafe.Add(mBase, uint32(v95))) = v99
			*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v99
			v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v102 + int32(1)
		} else {
			F_emscripten_builtin_free(m, v15)
			mBase = m.M
			v107 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v107 - v108
		}
		v113 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
		if v113 != v27 {
		} else {
			v119 = v16 + v27<<(uint(int32(3))%32) + int32(80)
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
			if v119 != v120 {
				v123 = v120
			} else {
				v123 = int32(0)
			}
			if v123 != 0 {
			} else {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
				if v124 != 0 {
					if v124 != v16+int32(88) {
						v139 = int32(1)
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
						if v131 != v16+int32(96) {
							v135 = int32(2)
						} else {
							v135 = int32(0)
						}
						if v131 != 0 {
							v137 = v135
						} else {
							v137 = int32(0)
						}
						v139 = v137
					}
				} else {
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
					if v131 != v16+int32(96) {
						v135 = int32(2)
					} else {
						v135 = int32(0)
					}
					if v131 != 0 {
						v137 = v135
					} else {
						v137 = int32(0)
					}
					v139 = v137
				}
				*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v139
			}
		}
	}
	return
}
