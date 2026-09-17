package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__crypt_gensalt_sha(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	if base.B2i32(l2 != v8)|base.B2i32(l4 <= int32(15)) == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
		v20 = l3 + int32(3)
		v23 = F_pg_snprintf(m, v20, int32(18), int32(_a_F__crypt_gensalt_sha_0), v9)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if v23 <= int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v165 = m.ExcPending
				if v165 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(2600))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F__crypt_gensalt_sha_1), int32(0))
						mBase = m.M
						v172 = m.ExcPending
						if v172 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F__crypt_gensalt_sha_2), int32(220), int32(_a_F__crypt_gensalt_sha_3))
							mBase = m.M
							v177 = m.ExcPending
							if v177 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v29 = v20 + v23
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v31 = int32(63)
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v35)
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)) = uint8(v42)
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)) = uint8(v49)
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)) = uint8(v56)
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)) = uint8(v63)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)) = uint8(v70)
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)) = uint8(v77)
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
				v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)) = uint8(v84)
				v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
				v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)) = uint8(v91)
				v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+9)) = uint8(v98)
				v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
				v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+10)) = uint8(v105)
				v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+11)) = uint8(v112)
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+12)) = uint8(v119)
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
				v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+13)) = uint8(v126)
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
				v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+14)) = uint8(v133)
				v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
				v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135&v31)+uint32(_c_F__crypt_gensalt_sha[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)) = uint8(v140)
				m.G0 = v9 + int32(16)
				return l3
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v149 = m.ExcPending
		if v149 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(2600))
			mBase = m.M
			v152 = m.ExcPending
			if v152 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F__crypt_gensalt_sha_4), int32(0))
				mBase = m.M
				v156 = m.ExcPending
				if v156 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F__crypt_gensalt_sha_2), int32(213), int32(_a_F__crypt_gensalt_sha_3))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F__crypt_gensalt_traditional_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	if base.B2i32(l2 < int32(2))|base.B2i32(l4 < int32(3)) == int32(0) {
		if base.B2i32(l0 == int32(0))|base.B2i32(l0 == int32(25)) != 0 {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v23 = int32(63)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22&v23)+uint32(_c_F__crypt_gensalt_traditional_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v27)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v30)
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29&v23)+uint32(_c_F__crypt_gensalt_traditional_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v36)
			return l3
		} else {
			v39 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v39)
			return v39
		}
	} else {
		if int32(0) < l4 {
			v39 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v39)
			return v39
		} else {
			return int32(0)
		}
	}
}
func F_run_crypt_bf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F__crypt_blowfish_rn(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
