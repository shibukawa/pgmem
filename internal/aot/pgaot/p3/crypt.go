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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	if l2 != v8 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v131 = m.ExcPending
		if v131 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(2600))
			mBase = m.M
			v134 = m.ExcPending
			if v134 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(222174), int32(0))
				mBase = m.M
				v140 = m.ExcPending
				if v140 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(484715), int32(213), int32(497839))
					mBase = m.M
					v147 = m.ExcPending
					if v147 != 0 {
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
		if l4 <= int32(15) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v131 = m.ExcPending
			if v131 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(222174), int32(0))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(484715), int32(213), int32(497839))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
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
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
			v17 = l3 + int32(3)
			v20 = F_pg_snprintf(m, v17, int32(18), int32(664409), v9)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v20 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(2600))
						mBase = m.M
						v154 = m.ExcPending
						if v154 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(324544), int32(0))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(484715), int32(220), int32(497839))
								mBase = m.M
								v167 = m.ExcPending
								if v167 != 0 {
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
					v26 = v17 + v20
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v29 = int32(63)
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v32)
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)) = uint8(v38)
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+2)) = uint8(v44)
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+3)) = uint8(v50)
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+4)) = uint8(v56)
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+5)) = uint8(v62)
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+6)) = uint8(v68)
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+7)) = uint8(v74)
					v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v80)
					v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
					v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+9)) = uint8(v86)
					v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+10)) = uint8(v92)
					v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+11)) = uint8(v98)
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
					v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+12)) = uint8(v104)
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
					v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+13)) = uint8(v110)
					v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
					v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+14)) = uint8(v116)
					v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
					v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118&v29)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v122)
					m.G0 = v9 + int32(16)
					return l3
				}
			}
		}
	}
}
func F__crypt_gensalt_traditional_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	if l2 < int32(2) {
		if int32(0) < l4 {
			v34 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v34)
			return v34
		} else {
			return int32(0)
		}
	} else {
		if l4 < int32(3) {
			if int32(0) < l4 {
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v34)
				return v34
			} else {
				return int32(0)
			}
		} else {
			if l0 == int32(0) {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v20 = int32(63)
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19&v20)+uint32(_consts[1297]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v23)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v26 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v26)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25&v20)+uint32(_consts[1297]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v31)
				return l3
			} else {
				if l0 == int32(25) {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v20 = int32(63)
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19&v20)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v23)
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v26 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v26)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25&v20)+uint32(_consts[1297]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v31)
					return l3
				} else {
					v34 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v34)
					return v34
				}
			}
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
