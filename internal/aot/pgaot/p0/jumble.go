package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleCreateRoleStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int64
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int64
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	v5 = l1 + int32(4)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v11 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v65 = v14
	} else {
		v15 = int32(4)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v17-int32(1021)) < base.Ui32(v15) {
			v26 = v17
			v28 = v15
			v30 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v26) {
					v35 = F_hash_bytes_extended(m, v16, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v16))) = v35
					v38 = int32(8)
				} else {
					v38 = v26
				}
				v40 = int32(1024) - v38
				if base.Ui32(v28) < base.Ui32(v40) {
					v42 = v28
				} else {
					v42 = v40
				}
				if v42 != 0 {
					base.MemoryCopy(m, v38+v16, v30, v42)
				} else {
				}
				v46 = v38 + v42
				v47 = v28 - v42
				if v47 != 0 {
					v26 = v46
					v28 = v47
					v30 = v42 + v30
					continue
				} else {
					break
				}
				break
			}
			v55 = v46
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17+v16))) = v11
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v55 = v50 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v55
		v65 = v55
	}
	v70 = int32(4)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v65-int32(1021)) < base.Ui32(v70) {
		v77 = v5
		v78 = v65
		v80 = v70
		for {
			if base.Ui32(int32(1024)) <= base.Ui32(v78) {
				v87 = F_hash_bytes_extended(m, v71, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v71))) = v87
				v90 = int32(8)
			} else {
				v90 = v78
			}
			v92 = int32(1024) - v90
			if base.Ui32(v80) < base.Ui32(v92) {
				v94 = v80
			} else {
				v94 = v92
			}
			if v94 != 0 {
				base.MemoryCopy(m, v90+v71, v77, v94)
			} else {
			}
			v98 = v90 + v94
			v99 = v80 - v94
			if v99 != 0 {
				v77 = v77 + v94
				v78 = v98
				v80 = v99
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v98
	} else {
		v102 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*int32)(unsafe.Add(mBase, uint32(v65+v71))) = v102
		v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v104 + int32(4)
	}
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v115 != 0 {
		v116 = F_strlen(m, v115)
		mBase = m.M
		v118 = v116 + int32(1)
		v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v124 == int32(0) {
			v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v181 = v127
		} else {
			v128 = int32(4)
			v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v130-int32(1021)) < base.Ui32(v128) {
				v140 = v130
				v142 = v128
				v144 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v140) {
						v149 = F_hash_bytes_extended(m, v129, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v129))) = v149
						v152 = int32(8)
					} else {
						v152 = v140
					}
					v154 = int32(1024) - v152
					if base.Ui32(v142) < base.Ui32(v154) {
						v156 = v142
					} else {
						v156 = v154
					}
					if v156 != 0 {
						base.MemoryCopy(m, v152+v129, v144, v156)
					} else {
					}
					v160 = v152 + v156
					v161 = v142 - v156
					if v161 != 0 {
						v140 = v160
						v142 = v161
						v144 = v156 + v144
						continue
					} else {
						break
					}
					break
				}
				v170 = v160
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v130+v129))) = v124
				v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v170 = v164 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170
			v181 = v170
		}
		v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(1024)-v181) < base.Ui32(v118) {
			v191 = v115
			v192 = v118
			v193 = v181
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v193) {
					v202 = F_hash_bytes_extended(m, v186, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v186))) = v202
					v205 = int32(8)
				} else {
					v205 = v193
				}
				v207 = int32(1024) - v205
				if base.Ui32(v192) < base.Ui32(v207) {
					v209 = v192
				} else {
					v209 = v207
				}
				if v209 != 0 {
					base.MemoryCopy(m, v205+v186, v191, v209)
				} else {
				}
				v213 = v205 + v209
				v214 = v192 - v209
				if v214 != 0 {
					v191 = v191 + v209
					v192 = v214
					v193 = v213
					continue
				} else {
					break
				}
				break
			}
			v222 = v213
		} else {
			if v118 != 0 {
				base.MemoryCopy(m, v181+v186, v115, v118)
			} else {
			}
			v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v222 = v217 + v118
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222
	} else {
		v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v228 + int32(1)
	}
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F__jumbleNode(m, l0, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		return
	} else {
		return
	}
}
