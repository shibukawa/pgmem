package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_subltree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	v4 = int32(0)
	if base.B2i32(l1|l2 < v4)|base.B2i32(l2 < l1) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v101 = m.ExcPending
		if v101 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_inner_subltree_0), int32(0))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_inner_subltree_1), int32(273), int32(_a_F_inner_subltree_2))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
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
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		if v18 <= l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_inner_subltree_0), int32(0))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_inner_subltree_1), int32(273), int32(_a_F_inner_subltree_2))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
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
			v21 = l0 + int32(8)
			if l2 < v18 {
				v23 = l2
			} else {
				v23 = v18
			}
			if v23 <= int32(0) {
				v174 = v21
				v183 = v21
			} else {
				v27 = v23 - int32(1)
				if v27 == int32(0) {
					v153 = v21
					v156 = v21
				} else {
					v30 = int32(3)
					v31 = v27 & v30
					if base.Ui32(v23-int32(2)) < base.Ui32(v30) {
						v116 = v21
						v118 = int32(0)
						v119 = v21
						v128 = v116
						v130 = v118
						v131 = v119
						v134 = v4
						for {
							v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128))))
							v145 = v128 + (v140+int32(9))&int32(_a_F_inner_subltree_3)
							v147 = v130 + int32(1)
							if v147 == l1 {
								v149 = v145
							} else {
								v149 = v131
							}
							v151 = v134 + int32(1)
							if v151 != v31 {
								v128 = v145
								v130 = v147
								v131 = v149
								v134 = v151
								continue
							} else {
								break
							}
							break
						}
						v153 = v145
						v156 = v149
					} else {
						v40 = v21
						v42 = int32(0)
						v43 = v21
						v51 = v4
						for {
							v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40))))
							v53 = int32(9)
							v55 = int32(_a_F_inner_subltree_3)
							v57 = v40 + (v52+v53)&v55
							v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57))))
							v63 = v57 + (v58+v53)&v55
							v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63))))
							v69 = v63 + (v64+v53)&v55
							v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69))))
							v75 = v69 + (v70+v53)&v55
							if v42|int32(1) == l1 {
								v79 = v57
							} else {
								v79 = v43
							}
							if v42|int32(2) == l1 {
								v83 = v63
							} else {
								v83 = v79
							}
							if v42|int32(3) == l1 {
								v87 = v69
							} else {
								v87 = v83
							}
							v89 = v42 + int32(4)
							if v89 == l1 {
								v91 = v75
							} else {
								v91 = v87
							}
							v93 = v51 + int32(4)
							if v93 != v27&int32(-4) {
								v40 = v75
								v42 = v89
								v43 = v91
								v51 = v93
								continue
							} else {
								break
							}
							break
						}
						if v31 == int32(0) {
							v153 = v75
							v156 = v91
						} else {
							v116 = v75
							v118 = v89
							v119 = v91
							v128 = v116
							v130 = v118
							v131 = v119
							v134 = v4
							for {
								v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128))))
								v145 = v128 + (v140+int32(9))&int32(_a_F_inner_subltree_3)
								v147 = v130 + int32(1)
								if v147 == l1 {
									v149 = v145
								} else {
									v149 = v131
								}
								v151 = v134 + int32(1)
								if v151 != v31 {
									v128 = v145
									v130 = v147
									v131 = v149
									v134 = v151
									continue
								} else {
									break
								}
								break
							}
							v153 = v145
							v156 = v149
						}
					}
				}
				v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153))))
				v174 = v156
				v183 = v153 + (v165+int32(9))&int32(_a_F_inner_subltree_3)
			}
			v184 = v183 - v174
			v186 = v184 + int32(8)
			v187 = F_palloc0(m, v186)
			mBase = m.M
			v188 = m.ExcPending
			if v188 != 0 {
				return int32(0)
			} else {
				v189 = v23 - l1
				*(*uint16)(unsafe.Add(mBase, uint32(v187)+4)) = uint16(v189)
				*(*int32)(unsafe.Add(mBase, uint32(v187))) = v186 << (uint(int32(2)) % 32)
				if v184 != 0 {
					base.MemoryCopy(m, v187+int32(8), v174, v184)
				} else {
				}
				return v187
			}
		}
	}
}
