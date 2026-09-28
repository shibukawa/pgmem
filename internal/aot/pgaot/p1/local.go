package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FlushLocalBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v111 int64
	_ = v111
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v178 int64
	_ = v178
	var v182 int32
	_ = v182
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v203 int64
	_ = v203
	var v208 int32
	_ = v208
	var v216 int64
	_ = v216
	var v219 int64
	_ = v219
	var v223 int32
	_ = v223
	var v225 int64
	_ = v225
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13+(int32(-2)-v15)<<(uint(int32(2))%32))))
	v22 = l0 + int32(36)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v23 != int32(-1) {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v26
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v28
		F_pgaio_wref_wait(m, v10+int32(32))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = int64(0)
			v37 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v34, v34)
			if v37&int64(8388608) != v34 {
				if l1 == int32(0) {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v46
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v48
					v50 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v50
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v48
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[1]))
					v57 = F_smgropen(m, v10+int32(8), v56)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						v59 = v57
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						F_PageSetChecksum(m, v20, v60)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[2])))
							v67 = m.G0
							v69 = v67 - int32(16)
							m.G0 = v69
							if v64 != 0 {
								F___clock_gettime(m, int32(1), v69)
								mBase = m.M
								v73 = int64(*(*int32)(unsafe.Add(mBase, uint32(v69)+8)))
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
								v78 = v73 + v74*int64(1000000000)
							} else {
								v78 = int64(0)
							}
							m.G0 = v69 + int32(16)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v20
							F_smgrwritev(m, v59, v83, v82, v10+int32(32), int32(0))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								v90 = int32(1)
								v94 = int64(8192)
								v98 = m.G0
								v100 = v98 - int32(16)
								m.G0 = v100
								if v78 != int64(0) {
									F___clock_gettime(m, int32(1), v100)
									mBase = m.M
									v106 = int64(*(*int32)(unsafe.Add(mBase, uint32(v100)+8)))
									v107 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
									v111 = v106 + (v107*int64(1000000000) - v78)
									v118 = int32(_a_F_FlushLocalBuffer_0)
									v120 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3]))
									v122 = base.I64_div_s(v111, int64(1000))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3])) = v120 + v122
									switch v90 {
									case 0:
										v125 = int32(_a_F_FlushLocalBuffer_1)
										v127 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4]))
										*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4])) = v127 + v111
									case 1:
										v130 = int32(_a_F_FlushLocalBuffer_2)
										v132 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5]))
										*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5])) = v132 + v111
									default:
									}
									v161 = int32(568)
									v162 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6])) = v162 + v111
									v166 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[7]))
									v173 = int32(0)
									if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v166))|base.B2i32(int32(1)<<(uint(v166)%32)&int32(_a_F_FlushLocalBuffer_3) == v173) == v173 {
										v178 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8]))
										*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8])) = v178 + v111
										v182 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v182)
										*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[10])) = uint8(v182)
									} else {
									}
								} else {
								}
								v198 = int32(568)
								v199 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11])) = v199 + base.I64_extend_i32_u(v90)
								v203 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12])) = v203 + v94
								F_pgstat_count_backend_io_op(m, v90, int32(3), int32(7), v90, v94)
								mBase = m.M
								v208 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v208)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[13])) = uint8(v208)
								m.G0 = v100 + int32(16)
								v216 = int64(0)
								v219 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v216, v216)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v219 & int64(-142606337)
								v223 = int32(_a_F_FlushLocalBuffer_4)
								v225 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14])) = v225 + int64(1)
								m.G0 = v10 + int32(48)
								return
							}
						}
					}
				} else {
					v59 = l1
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					F_PageSetChecksum(m, v20, v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[2])))
						v67 = m.G0
						v69 = v67 - int32(16)
						m.G0 = v69
						if v64 != 0 {
							F___clock_gettime(m, int32(1), v69)
							mBase = m.M
							v73 = int64(*(*int32)(unsafe.Add(mBase, uint32(v69)+8)))
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
							v78 = v73 + v74*int64(1000000000)
						} else {
							v78 = int64(0)
						}
						m.G0 = v69 + int32(16)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v20
						F_smgrwritev(m, v59, v83, v82, v10+int32(32), int32(0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							v90 = int32(1)
							v94 = int64(8192)
							v98 = m.G0
							v100 = v98 - int32(16)
							m.G0 = v100
							if v78 != int64(0) {
								F___clock_gettime(m, int32(1), v100)
								mBase = m.M
								v106 = int64(*(*int32)(unsafe.Add(mBase, uint32(v100)+8)))
								v107 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
								v111 = v106 + (v107*int64(1000000000) - v78)
								v118 = int32(_a_F_FlushLocalBuffer_0)
								v120 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3]))
								v122 = base.I64_div_s(v111, int64(1000))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3])) = v120 + v122
								switch v90 {
								case 0:
									v125 = int32(_a_F_FlushLocalBuffer_1)
									v127 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4])) = v127 + v111
								case 1:
									v130 = int32(_a_F_FlushLocalBuffer_2)
									v132 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5])) = v132 + v111
								default:
								}
								v161 = int32(568)
								v162 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6])) = v162 + v111
								v166 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[7]))
								v173 = int32(0)
								if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v166))|base.B2i32(int32(1)<<(uint(v166)%32)&int32(_a_F_FlushLocalBuffer_3) == v173) == v173 {
									v178 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8])) = v178 + v111
									v182 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v182)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[10])) = uint8(v182)
								} else {
								}
							} else {
							}
							v198 = int32(568)
							v199 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11])) = v199 + base.I64_extend_i32_u(v90)
							v203 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12])) = v203 + v94
							F_pgstat_count_backend_io_op(m, v90, int32(3), int32(7), v90, v94)
							mBase = m.M
							v208 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v208)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[13])) = uint8(v208)
							m.G0 = v100 + int32(16)
							v216 = int64(0)
							v219 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v216, v216)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v219 & int64(-142606337)
							v223 = int32(_a_F_FlushLocalBuffer_4)
							v225 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14])) = v225 + int64(1)
							m.G0 = v10 + int32(48)
							return
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v235 = m.ExcPending
				if v235 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_FlushLocalBuffer_5), int32(0))
					mBase = m.M
					v239 = m.ExcPending
					if v239 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_FlushLocalBuffer_6), int32(196), int32(_a_F_FlushLocalBuffer_7))
						mBase = m.M
						v244 = m.ExcPending
						if v244 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v34 = int64(0)
		v37 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v34, v34)
		if v37&int64(8388608) != v34 {
			if l1 == int32(0) {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v44
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v46
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v48
				v50 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v50
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v48
				v56 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[1]))
				v57 = F_smgropen(m, v10+int32(8), v56)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v59 = v57
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					F_PageSetChecksum(m, v20, v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[2])))
						v67 = m.G0
						v69 = v67 - int32(16)
						m.G0 = v69
						if v64 != 0 {
							F___clock_gettime(m, int32(1), v69)
							mBase = m.M
							v73 = int64(*(*int32)(unsafe.Add(mBase, uint32(v69)+8)))
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
							v78 = v73 + v74*int64(1000000000)
						} else {
							v78 = int64(0)
						}
						m.G0 = v69 + int32(16)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v20
						F_smgrwritev(m, v59, v83, v82, v10+int32(32), int32(0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							v90 = int32(1)
							v94 = int64(8192)
							v98 = m.G0
							v100 = v98 - int32(16)
							m.G0 = v100
							if v78 != int64(0) {
								F___clock_gettime(m, int32(1), v100)
								mBase = m.M
								v106 = int64(*(*int32)(unsafe.Add(mBase, uint32(v100)+8)))
								v107 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
								v111 = v106 + (v107*int64(1000000000) - v78)
								v118 = int32(_a_F_FlushLocalBuffer_0)
								v120 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3]))
								v122 = base.I64_div_s(v111, int64(1000))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3])) = v120 + v122
								switch v90 {
								case 0:
									v125 = int32(_a_F_FlushLocalBuffer_1)
									v127 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4])) = v127 + v111
								case 1:
									v130 = int32(_a_F_FlushLocalBuffer_2)
									v132 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5])) = v132 + v111
								default:
								}
								v161 = int32(568)
								v162 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6])) = v162 + v111
								v166 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[7]))
								v173 = int32(0)
								if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v166))|base.B2i32(int32(1)<<(uint(v166)%32)&int32(_a_F_FlushLocalBuffer_3) == v173) == v173 {
									v178 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8])) = v178 + v111
									v182 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v182)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[10])) = uint8(v182)
								} else {
								}
							} else {
							}
							v198 = int32(568)
							v199 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11])) = v199 + base.I64_extend_i32_u(v90)
							v203 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12])) = v203 + v94
							F_pgstat_count_backend_io_op(m, v90, int32(3), int32(7), v90, v94)
							mBase = m.M
							v208 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v208)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[13])) = uint8(v208)
							m.G0 = v100 + int32(16)
							v216 = int64(0)
							v219 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v216, v216)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v219 & int64(-142606337)
							v223 = int32(_a_F_FlushLocalBuffer_4)
							v225 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14])) = v225 + int64(1)
							m.G0 = v10 + int32(48)
							return
						}
					}
				}
			} else {
				v59 = l1
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				F_PageSetChecksum(m, v20, v60)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[2])))
					v67 = m.G0
					v69 = v67 - int32(16)
					m.G0 = v69
					if v64 != 0 {
						F___clock_gettime(m, int32(1), v69)
						mBase = m.M
						v73 = int64(*(*int32)(unsafe.Add(mBase, uint32(v69)+8)))
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
						v78 = v73 + v74*int64(1000000000)
					} else {
						v78 = int64(0)
					}
					m.G0 = v69 + int32(16)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v20
					F_smgrwritev(m, v59, v83, v82, v10+int32(32), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						v90 = int32(1)
						v94 = int64(8192)
						v98 = m.G0
						v100 = v98 - int32(16)
						m.G0 = v100
						if v78 != int64(0) {
							F___clock_gettime(m, int32(1), v100)
							mBase = m.M
							v106 = int64(*(*int32)(unsafe.Add(mBase, uint32(v100)+8)))
							v107 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
							v111 = v106 + (v107*int64(1000000000) - v78)
							v118 = int32(_a_F_FlushLocalBuffer_0)
							v120 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3]))
							v122 = base.I64_div_s(v111, int64(1000))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3])) = v120 + v122
							switch v90 {
							case 0:
								v125 = int32(_a_F_FlushLocalBuffer_1)
								v127 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4])) = v127 + v111
							case 1:
								v130 = int32(_a_F_FlushLocalBuffer_2)
								v132 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5])) = v132 + v111
							default:
							}
							v161 = int32(568)
							v162 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6])) = v162 + v111
							v166 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[7]))
							v173 = int32(0)
							if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v166))|base.B2i32(int32(1)<<(uint(v166)%32)&int32(_a_F_FlushLocalBuffer_3) == v173) == v173 {
								v178 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8])) = v178 + v111
								v182 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v182)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[10])) = uint8(v182)
							} else {
							}
						} else {
						}
						v198 = int32(568)
						v199 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11])) = v199 + base.I64_extend_i32_u(v90)
						v203 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12])) = v203 + v94
						F_pgstat_count_backend_io_op(m, v90, int32(3), int32(7), v90, v94)
						mBase = m.M
						v208 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v208)
						*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[13])) = uint8(v208)
						m.G0 = v100 + int32(16)
						v216 = int64(0)
						v219 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v216, v216)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v219 & int64(-142606337)
						v223 = int32(_a_F_FlushLocalBuffer_4)
						v225 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14])) = v225 + int64(1)
						m.G0 = v10 + int32(48)
						return
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v235 = m.ExcPending
			if v235 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_FlushLocalBuffer_5), int32(0))
				mBase = m.M
				v239 = m.ExcPending
				if v239 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_FlushLocalBuffer_6), int32(196), int32(_a_F_FlushLocalBuffer_7))
					mBase = m.M
					v244 = m.ExcPending
					if v244 != 0 {
						return
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
func F_LocalExecuteInvalidationMessage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	var v310 int64
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v358 int32
	_ = v358
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v685 int32
	_ = v685
	var v686 int64
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v739 int32
	_ = v739
	var v740 int64
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int64
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v776 int64
	_ = v776
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if v2 <= v15 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L10
	} else {
		goto L250
	}
L2:
	;
	m.G0 = v13 - int32(-64)
	return
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	switch v15&int32(255) - int32(250) {
	case 0:
		goto L74
	case 1:
		goto L75
	case 2:
		goto L76
	case 3:
		goto L72
	case 4:
		goto L77
	case 5:
		goto L78
	default:
		goto L73
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v18 != v20 {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	return
L11:
	;
	v24 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	if base.Ui32(v24) < base.Ui32(int32(85)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v214 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(int32(85)) <= base.Ui32(v214) {
		goto L1
	} else {
		goto L66
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_c_F_LocalExecuteInvalidationMessage[1])))
	if v34 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L63
	}
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if int32(0) < v35 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	m.G0 = v28 + int32(16)
	goto L12
L19:
	;
	v40 = v35
	v44 = v2
	goto L22
L20:
	;
	goto L21
L21:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v109 = v102 + (v103-int32(1))&v25<<(uint(int32(3))%32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v111 = int32(0)
	if base.B2i32(v110 == v111)|base.B2i32(v110 == v109) == v111 {
		goto L36
	} else {
		goto L37
	}
L22:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
	v51 = v48 + v44<<(uint(int32(3))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v53 = int32(0)
	if base.B2i32(v52 == v53)|base.B2i32(v52 == v51) == v53 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v60 = v52
	goto L27
L25:
	;
	v81 = v40
	goto L26
L26:
	;
	v90 = v44 + int32(1)
	if v90 < v81 {
		v40 = v81
		v44 = v90
		goto L22
	} else {
		goto L35
	}
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+48))
	if int32(0) < v70 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v81 = v78
	goto L26
L29:
	;
	if v51 != v69 {
		v60 = v69
		goto L27
	} else {
		goto L34
	}
L30:
	;
	v73 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+52)) = uint8(v73)
	goto L29
L31:
	;
	goto L32
L32:
	;
	F_CatCacheRemoveCList(m, v34, v60)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	goto L28
L35:
	;
	goto L23
L36:
	;
	v119 = v110
	goto L39
L37:
	;
	goto L38
L38:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[2]))
	if v157 != 0 {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	if v25 != v128 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	if v127 != v109 {
		v119 = v127
		goto L39
	} else {
		goto L50
	}
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v119)+48))
	if v130 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	F_CatCacheRemoveCTup(m, v34, v119)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L10
	} else {
		goto L49
	}
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v119)+76))
	if v133 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v140 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+52)) = uint8(v140)
	goto L41
L47:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	if v136 <= int32(0) {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	goto L41
L50:
	;
	goto L40
L51:
	;
	v159 = v157
	goto L54
L52:
	;
	goto L53
L53:
	;
	goto L18
L54:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v168 != v34 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	if v177 != 0 {
		v159 = v177
		goto L54
	} else {
		goto L62
	}
L57:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+8)))
	if v170 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v173 != v25 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+9)) = uint8(v175)
	goto L56
L61:
	;
	goto L60
L62:
	;
	goto L55
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v24
	F_errmsg_internal(m, int32(_a_F_LocalExecuteInvalidationMessage_0), v28)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_LocalExecuteInvalidationMessage_1), int32(694), int32(_a_F_LocalExecuteInvalidationMessage_2))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v219 = int32(*(*int16)(unsafe.Add(mBase, uint32(v214<<(uint(int32(1))%32))+uint32(_c_F_LocalExecuteInvalidationMessage[3]))))
	if v219 <= int32(0) {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v224 = v219
	goto L68
L68:
	;
	v236 = v224 << (uint(int32(4)) % 32) & int32(_a_F_LocalExecuteInvalidationMessage_3)
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v236)+uint32(_c_F_LocalExecuteInvalidationMessage[4])))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v236)+uint32(_c_F_LocalExecuteInvalidationMessage[5])))
	m.T0[v242].(func(*base.Module, int64, int32, int32))(m, v239, v214, v222)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L10
	} else {
		goto L70
	}
L69:
	;
	goto L2
L70:
	;
	v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(v236)+uint32(_c_F_LocalExecuteInvalidationMessage[6]))))
	if int32(0) < v247 {
		v224 = v247
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v765
	v767 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v767
	v769 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v770 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v769 | v770<<(uint(int32(16))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v767
	v776 = *(*int64)(unsafe.Add(mBase, uint32(v13)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v776
	v781 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[7]))
	if v781 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L73:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L10
	} else {
		goto L239
	}
L74:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v719 != v721 {
		goto L2
	} else {
		goto L233
	}
L75:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v708 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L76:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v696 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L77:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v369 != 0 {
		goto L108
	} else {
		goto L109
	}
L78:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v254 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v254 != v256 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L10
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[8]))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	if v263 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v265 = v263
	goto L87
L85:
	;
	goto L86
L86:
	;
	goto L2
L87:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v265-int32(12))))
	if v260 == v276 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v279 = v265 - int32(100)
	F_ResetCatalogCache(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L10
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	if v358 != 0 {
		v265 = v358
		goto L87
	} else {
		goto L107
	}
L92:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v283 = m.G0
	v285 = v283 - int32(16)
	m.G0 = v285
	if base.Ui32(v282) < base.Ui32(int32(85)) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L91
L94:
	;
	v291 = int32(*(*int16)(unsafe.Add(mBase, uint32(v282<<(uint(int32(1))%32))+uint32(_c_F_LocalExecuteInvalidationMessage[3]))))
	if int32(0) < v291 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L10
	} else {
		goto L104
	}
L97:
	;
	v294 = v291
	goto L100
L98:
	;
	goto L99
L99:
	;
	m.G0 = v285 + int32(16)
	goto L93
L100:
	;
	v307 = v294 << (uint(int32(4)) % 32) & int32(_a_F_LocalExecuteInvalidationMessage_3)
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v307)+uint32(_c_F_LocalExecuteInvalidationMessage[4])))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v307)+uint32(_c_F_LocalExecuteInvalidationMessage[5])))
	m.T0[v314].(func(*base.Module, int64, int32, int32))(m, v310, v282, int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L10
	} else {
		goto L102
	}
L101:
	;
	goto L99
L102:
	;
	v319 = int32(*(*int16)(unsafe.Add(mBase, uint32(v307)+uint32(_c_F_LocalExecuteInvalidationMessage[6]))))
	if int32(0) < v319 {
		v294 = v319
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v282
	F_errmsg_internal(m, int32(_a_F_LocalExecuteInvalidationMessage_0), v285)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_LocalExecuteInvalidationMessage_4), int32(1900), int32(_a_F_LocalExecuteInvalidationMessage_5))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	goto L88
L108:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v369 != v371 {
		goto L2
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v373 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L110
L112:
	;
	v669 = int32(0)
	v671 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[9]))
	if v671 <= v669 {
		goto L2
	} else {
		goto L216
	}
L113:
	;
	F_RelationCacheInvalidate(m)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L10
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v378 = m.G0
	v380 = v378 - int32(16)
	m.G0 = v380
	*(*int32)(unsafe.Add(mBase, uint32(v380)+12)) = v373
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[10]))
	v387 = int32(0)
	v389 = F_hash_search(m, v384, v380+int32(12), v387, v387)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L10
	} else {
		goto L119
	}
L116:
	;
	goto L112
L117:
	;
	m.G0 = v380 + int32(16)
	goto L112
L118:
	;
	v454 = int32(_a_F_LocalExecuteInvalidationMessage_6)
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[11])) = v456 + int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	if v460 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L119:
	;
	if v389 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v391 != 0 {
		goto L118
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[12]))
	if v394 <= int32(0) {
		goto L117
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v397 = int32(0)
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[13]))
	if v394 != int32(1) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v407 = v397
	v412 = v2
	goto L128
L126:
	;
	v437 = v397
	goto L127
L127:
	;
	v448 = v399 + v437<<(uint(int32(3))%32)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	if v449 != v450 {
		goto L117
	} else {
		goto L138
	}
L128:
	;
	v418 = v399 + v407<<(uint(int32(3))%32)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	if v419 == v420 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v394&int32(1) == int32(0) {
		goto L117
	} else {
		goto L137
	}
L130:
	;
	v422 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+4)) = uint8(v422)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	if v424 == v425 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v427 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+12)) = uint8(v427)
	goto L135
L134:
	;
	goto L135
L135:
	;
	v429 = int32(2)
	v430 = v407 + v429
	v432 = v412 + v429
	if v432 != v394&int32(2147483646) {
		v407 = v430
		v412 = v432
		goto L128
	} else {
		goto L136
	}
L136:
	;
	goto L129
L137:
	;
	v437 = v430
	goto L127
L138:
	;
	v452 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v448)+4)) = uint8(v452)
	goto L117
L139:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	if v546 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L140:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v391)+40))
	if v463 == int32(0) {
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[14]))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+20))
	goto L145
L143:
	;
	goto L142
L144:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	if v506 != 0 {
		goto L156
	} else {
		goto L157
	}
L145:
	;
	if base.B2i32(v468 == int32(2)) == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v391)+44))
	if v473 != 0 {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[15]))
	F_ResourceOwnerEnlarge(m, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v391)+16)) = v478 + int32(1)
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[16]))
	if v483 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[15]))
	F_ResourceOwnerRemember(m, v485, base.I64_extend_i32_u(v391), int32(_a_F_LocalExecuteInvalidationMessage_7))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L10
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	F_RelationRebuildRelation(m, v391)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L10
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v391)+16)) = v492 - int32(1)
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[16]))
	if v497 == int32(0) {
		goto L117
	} else {
		goto L154
	}
L154:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[15]))
	F_ResourceOwnerForget(m, v501, base.I64_extend_i32_u(v391), int32(_a_F_LocalExecuteInvalidationMessage_7))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	goto L117
L156:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v506)+72))
	v511 = v509 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+72)) = v511
	if v511 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	goto L158
L158:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v391)+256))
	if v539 != 0 {
		goto L168
	} else {
		goto L169
	}
L159:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	F_smgrclose(m, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L10
	} else {
		goto L167
	}
L160:
	;
	v516 = v506 + int32(76)
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17]))
	if v518 != 0 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	goto L162
L162:
	;
	goto L159
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+76)) = v525
	v527 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+80)) = v527
	*(*int32)(unsafe.Add(mBase, uint32(v525)+4)) = v516
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18])) = v516
	goto L162
L164:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18]))
	v525 = v520
	goto L163
L165:
	;
	goto L166
L166:
	;
	v522 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17])) = v522
	v525 = v522
	goto L163
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+12)) = int32(0)
	goto L158
L168:
	;
	F_pfree(m, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L10
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v542 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v391)+26)) = uint8(v542)
	*(*int32)(unsafe.Add(mBase, uint32(v391)+256)) = v542
	goto L117
L171:
	;
	goto L170
L172:
	;
	F_RelationClearRelation(m, v391)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L10
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[14]))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+20))
	goto L176
L175:
	;
	goto L117
L176:
	;
	if base.B2i32(v553 == int32(2)) == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	if v558 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	goto L179
L179:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+25)))
	if v598 != int32(1) {
		goto L196
	} else {
		goto L197
	}
L180:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v558)+72))
	v563 = v561 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v558)+72)) = v563
	if v563 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	goto L182
L182:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v391)+256))
	if v591 != 0 {
		goto L192
	} else {
		goto L193
	}
L183:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	F_smgrclose(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L10
	} else {
		goto L191
	}
L184:
	;
	v568 = v558 + int32(76)
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17]))
	if v570 != 0 {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	goto L186
L186:
	;
	goto L183
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v558)+76)) = v577
	v579 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, uint32(v558)+80)) = v579
	*(*int32)(unsafe.Add(mBase, uint32(v577)+4)) = v568
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18])) = v568
	goto L186
L188:
	;
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18]))
	v577 = v572
	goto L187
L189:
	;
	goto L190
L190:
	;
	v574 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17])) = v574
	v577 = v574
	goto L187
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+12)) = int32(0)
	goto L182
L192:
	;
	F_pfree(m, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L10
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v594 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v391)+26)) = uint8(v594)
	*(*int32)(unsafe.Add(mBase, uint32(v391)+256)) = v594
	goto L117
L195:
	;
	goto L194
L196:
	;
	F_RelationRebuildRelation(m, v391)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L10
	} else {
		goto L215
	}
L197:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	if v601 != int32(1) {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	if v604 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v604)+72))
	v609 = v607 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v604)+72)) = v609
	if v609 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	goto L201
L201:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v391)+256))
	if v637 != 0 {
		goto L211
	} else {
		goto L212
	}
L202:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	F_smgrclose(m, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L10
	} else {
		goto L210
	}
L203:
	;
	v614 = v604 + int32(76)
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17]))
	if v616 != 0 {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	goto L205
L205:
	;
	goto L202
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v604)+76)) = v623
	v625 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, uint32(v604)+80)) = v625
	*(*int32)(unsafe.Add(mBase, uint32(v623)+4)) = v614
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18])) = v614
	goto L205
L207:
	;
	v618 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18]))
	v623 = v618
	goto L206
L208:
	;
	goto L209
L209:
	;
	v620 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17])) = v620
	v623 = v620
	goto L206
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+12)) = int32(0)
	goto L201
L211:
	;
	F_pfree(m, v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L10
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v640 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v391)+26)) = uint8(v640)
	*(*int32)(unsafe.Add(mBase, uint32(v391)+256)) = v640
	goto L117
L214:
	;
	goto L213
L215:
	;
	goto L117
L216:
	;
	v675 = v669
	goto L217
L217:
	;
	v685 = v675 << (uint(int32(4)) % 32)
	v686 = *(*int64)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_LocalExecuteInvalidationMessage[19])))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_LocalExecuteInvalidationMessage[20])))
	m.T0[v688].(func(*base.Module, int64, int32))(m, v686, v687)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L10
	} else {
		goto L219
	}
L218:
	;
	goto L2
L219:
	;
	v692 = v675 + int32(1)
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[9]))
	if v692 < v694 {
		v675 = v692
		goto L217
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	F_RelationMapInvalidate(m, int32(1))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L10
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v696 != v703 {
		goto L2
	} else {
		goto L225
	}
L224:
	;
	goto L2
L225:
	;
	F_RelationMapInvalidate(m, int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L10
	} else {
		goto L226
	}
L226:
	;
	goto L2
L227:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L10
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v708 != v714 {
		goto L2
	} else {
		goto L231
	}
L230:
	;
	goto L2
L231:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	goto L2
L233:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[21]))
	if v724 <= int32(0) {
		goto L2
	} else {
		goto L234
	}
L234:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v729 = int32(0)
	goto L235
L235:
	;
	v739 = v729 << (uint(int32(4)) % 32)
	v740 = *(*int64)(unsafe.Add(mBase, uint32(v739)+uint32(_c_F_LocalExecuteInvalidationMessage[22])))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v739)+uint32(_c_F_LocalExecuteInvalidationMessage[23])))
	m.T0[v741].(func(*base.Module, int64, int32))(m, v740, v727)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L10
	} else {
		goto L237
	}
L236:
	;
	goto L2
L237:
	;
	v745 = v729 + int32(1)
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[21]))
	if v745 < v747 {
		v729 = v745
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	v753 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v753
	F_errmsg_internal(m, int32(_a_F_LocalExecuteInvalidationMessage_9), v11+int32(-48))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L10
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_LocalExecuteInvalidationMessage_4), int32(901), int32(_a_F_LocalExecuteInvalidationMessage_10))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L10
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	goto L2
L243:
	;
	v784 = int32(0)
	v786 = F_hash_search(m, v781, v11+int32(-32), v784, v784)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L10
	} else {
		goto L244
	}
L244:
	;
	if v786 == int32(0) {
		goto L242
	} else {
		goto L245
	}
L245:
	;
	v790 = int32(_a_F_LocalExecuteInvalidationMessage_11)
	v792 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[24])) = v792 + int32(1)
	F_mdclose(m, v786, int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L10
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v786)+20)) = int32(-1)
	F_mdclose(m, v786, int32(1))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L10
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v786)+24)) = int32(-1)
	F_mdclose(m, v786, int32(2))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L10
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v786)+28)) = int32(-1)
	F_mdclose(m, v786, int32(3))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L10
	} else {
		goto L249
	}
L249:
	;
	v814 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v786)+16)) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v786)+32)) = v814
	v818 = int32(_a_F_LocalExecuteInvalidationMessage_11)
	v820 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[24])) = v820 - int32(1)
	goto L242
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v214
	F_errmsg_internal(m, int32(_a_F_LocalExecuteInvalidationMessage_0), v13)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L10
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_LocalExecuteInvalidationMessage_4), int32(1900), int32(_a_F_LocalExecuteInvalidationMessage_5))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L10
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_add_local_real_reloption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64, l5 float64, l6 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v11 = F_palloc(m, int32(48))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = F_pstrdup(m, l1)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
			if l2 != 0 {
				v16 = F_pstrdup(m, l2)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = v16
					v19 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v19
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v18
					v22 = F_strlen(m, l1)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v22
					*(*float64)(unsafe.Add(mBase, uint32(v11)+40)) = l5
					*(*float64)(unsafe.Add(mBase, uint32(v11)+32)) = l4
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
					v32 = F_palloc(m, int32(8))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = l6
						*(*int32)(unsafe.Add(mBase, uint32(v32))) = v11
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v37 = F_lappend(m, v36, v32)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v37
							return
						}
					}
				}
			} else {
				v18 = int32(0)
				v19 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v18
				v22 = F_strlen(m, l1)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v22
				*(*float64)(unsafe.Add(mBase, uint32(v11)+40)) = l5
				*(*float64)(unsafe.Add(mBase, uint32(v11)+32)) = l4
				*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
				v32 = F_palloc(m, int32(8))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = l6
					*(*int32)(unsafe.Add(mBase, uint32(v32))) = v11
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v37 = F_lappend(m, v36, v32)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v37
						return
					}
				}
			}
		}
	}
}
func F_local_buffer_readv_complete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v180 int64
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int64
	_ = v195
	var v197 int64
	_ = v197
	var v201 int64
	_ = v201
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int64
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	v5 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(32)
	m.G0 = v29
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v31
	v35 = base.I32_wrap_i64(v31) & int32(448)
	v37 = l1 + int32(104)
	goto L1
L1:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(23)))) = uint8(v40)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_complete[0]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	goto L2
L2:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+23)))
	if v49 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v58 = l3 & int32(1)
	if v58 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v247 = v5
	v248 = v5
	v251 = v5
	v253 = v5
	v254 = v5
	v255 = v5
	v266 = int32(0)
	goto L5
L5:
	;
	if v35 == int32(256) {
		goto L55
	} else {
		goto L56
	}
L6:
	;
	v59 = int32(10)
	goto L8
L7:
	;
	v59 = int32(2)
	goto L8
L8:
	;
	v65 = int32(0)
	v71 = v5
	v72 = v5
	v75 = v5
	v76 = v5
	v77 = v5
	v78 = v5
	v79 = v5
	goto L9
L9:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_complete[1]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v44+v45<<(uint(int32(3))%32)+v65<<(uint(int32(3))%32))))
	v99 = v97 ^ int32(-1)
	v102 = v93 + v99*int32(56)
	if v97 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v247 = v225
	v248 = v228
	v251 = v215
	v253 = v219
	v254 = v223
	v255 = v230
	v266 = base.B2i32(v224&int32(255) != int32(0))
	goto L5
L11:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)) = uint8(v120)
	if base.B2i32(v35 == int32(256))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(v31)>>(uint(int64(32))%64))) <= v65) != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_complete[2]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v99<<(uint(int32(2))%32))))
	v118 = v110
	goto L11
L13:
	;
	goto L14
L14:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_complete[3]))
	v118 = v112 + v97<<(uint(int32(13))%32) + int32(-8192)
	goto L11
L15:
	;
	v197 = int64(0)
	v201 = base.AtomicRmwCmpxchg64(m, v102, int32(24), v197, v197)
	goto L33
L16:
	;
	v188 = int32(1)
	v189 = v120
	v190 = int32(0)
	v195 = int64(134217728)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v128 = F_PageIsVerified(m, v118, v119, l3&int32(4)|v59, v29+int32(22))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
	if v128 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v161<<(uint(int32(18))%32) | v165 | (v163 | (v160<<(uint(int32(9))%32) | v162)) | v65<<(uint(int32(25))%32)
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v29)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v180
	F_pgaio_result_report(m, v29+int32(8), v37, int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L19
	} else {
		goto L31
	}
L22:
	;
	v133 = int32(2048)
	if v58 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v148 = int32(1)
	v149 = int64(16777216)
	if v130&v148 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v137 = int32(0)
	v159 = v137
	v160 = v120
	v161 = v130
	v162 = v137
	v163 = v133
	v164 = int64(134217728)
	v165 = int32(259)
	goto L21
L26:
	;
	goto L27
L27:
	;
	v140 = int32(0)
	base.MemoryFill(m, v118, v140, int32(_a_F_local_buffer_readv_complete_0))
	v145 = int32(1)
	v159 = v145
	v160 = v145
	v161 = v130
	v162 = v140
	v163 = v133
	v164 = int64(16777216)
	v165 = int32(195)
	goto L21
L28:
	;
	v188 = v148
	v189 = v120
	v190 = int32(0)
	v195 = v149
	goto L15
L29:
	;
	goto L30
L30:
	;
	v159 = v148
	v160 = v120
	v161 = int32(1)
	v162 = int32(1024)
	v163 = int32(0)
	v164 = v149
	v165 = int32(195)
	goto L21
L31:
	;
	v188 = v159 | v160
	v189 = v160
	v190 = v128
	v195 = v164
	goto L15
L32:
	;
	if v76&int32(255) != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	F_pgaio_wref_clear(m, v102+int32(36))
	mBase = m.M
	goto L35
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v102)+24)) = v201&int64(-134217729) - int64(1) | v195
	goto L32
L36:
	;
	v214 = v75
	goto L38
L37:
	;
	v214 = v65
	goto L38
L38:
	;
	if v190 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v215 = v214
	goto L41
L40:
	;
	v215 = v75
	goto L41
L41:
	;
	if v71&int32(255) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v218 = v77
	goto L44
L43:
	;
	v218 = v65
	goto L44
L44:
	;
	if v189 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v219 = v218
	goto L47
L46:
	;
	v219 = v77
	goto L47
L47:
	;
	if v72&int32(255) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v222 = v78
	goto L50
L49:
	;
	v222 = v65
	goto L50
L50:
	;
	if v188 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v223 = v78
	goto L53
L52:
	;
	v223 = v222
	goto L53
L53:
	;
	v224 = v190 + v76
	v225 = v189 + v71
	v226 = int32(1)
	v228 = v72 + (v188 ^ v226)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
	v230 = v229 + v79
	v232 = v65 + v226
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+23)))
	if base.Ui32(v232) < base.Ui32(v233) {
		v65 = v232
		v71 = v225
		v72 = v228
		v75 = v215
		v76 = v224
		v77 = v219
		v78 = v223
		v79 = v230
		goto L9
	} else {
		goto L54
	}
L54:
	;
	goto L10
L55:
	;
	v317 = v255 & int32(255)
	if v317 != 0 {
		goto L77
	} else {
		goto L78
	}
L56:
	;
	v269 = int32(255)
	v270 = v247 & v269
	v272 = v248 & v269
	v273 = int32(0)
	if v270|(base.B2i32(v272 != v273)|v266)&int32(1) == v273 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	if v272 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v287 = int32(259)
	goto L60
L59:
	;
	v287 = int32(195)
	goto L60
L60:
	;
	if v270 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v291 = int32(512)
	goto L63
L62:
	;
	v291 = int32(0)
	goto L63
L63:
	;
	if v266 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v294 = int32(1024)
	goto L66
L65:
	;
	v294 = int32(0)
	goto L66
L66:
	;
	if v272 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v296 = v248
	goto L69
L68:
	;
	v296 = v247
	goto L69
L69:
	;
	if v270 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v303 = v253
	goto L72
L71:
	;
	v303 = v251
	goto L72
L72:
	;
	if v272 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v304 = v254
	goto L75
L74:
	;
	v304 = v303
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v255&int32(255)<<(uint(int32(18))%32) | v287 | (v291 | v294 | v296&int32(255)<<(uint(int32(11))%32)) | v304<<(uint(int32(25))%32)
	v309 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v309
	F_pgaio_result_report(m, v29, v37, int32(14))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L19
	} else {
		goto L76
	}
L76:
	;
	goto L55
L77:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	F_pgstat_report_checksum_failures_in_db(m, v318, v317)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L19
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	m.G0 = v29 + int32(32)
	return
L80:
	;
	goto L79
}
func F_local_ts_extend_down(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v3 = l2
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = F_MemoryContextAlloc(m, v7, int32(24))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(1024)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	v17 = l3 - int32(8)
	if int32(0) < v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = v9
	v26 = base.I64_extend_i32_u(v17)
	goto L6
L4:
	;
	v46 = v9
	goto L5
L5:
	;
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)) = uint8(v48)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+3)) = uint8(v3)
	return v46 + int32(8)
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = F_MemoryContextAlloc(m, v27, int32(24))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v46 = v29
	goto L5
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v29
	v34 = int64(base.Ui64(v3) >> (uint(v26) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)) = uint8(v34)
	v36 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)) = uint8(v36)
	v38 = int64(8)
	if base.Ui64(v38) < base.Ui64(v26) {
		v25 = v29
		v26 = v26 - v38
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_read_local_xlog_page_guts(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = l1 + base.I64_extend_i32_s(l2)
	goto L2
L1:
	;
	if base.Ui64(v103) < base.Ui64(l1-int64(-8192)) {
		goto L33
	} else {
		goto L34
	}
L2:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[0])))
	if v29 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
	v102 = v86
	v103 = v101
	goto L1
L4:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_XLogReadDetermineTimeline(m, l0, l1, l2, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L16
	} else {
		goto L20
	}
L5:
	;
	if v39 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[1]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+308))
	v37 = base.B2i32(v35 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[0])) = uint8(v37)
	v39 = v37
	goto L8
L7:
	;
	v39 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v43 = v13 + int32(4)
	v45 = int32(_a_F_read_local_xlog_page_guts_0)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[1]))
	v47 = int64(0)
	v50 = base.AtomicRmwCmpxchg64(m, v46, int32(272), v47, v47)
	*(*int64)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[2])) = v50
	v52 = int32(0)
	v55 = base.AtomicRmwOr32(m, v52, int32(_a_F_read_local_xlog_page_guts_1), v52)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[1]))
	v62 = base.AtomicRmwCmpxchg64(m, v58, int32(264), v47, v47)
	*(*int64)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[3])) = v62
	if v43 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v72 = F_GetXLogReplayRecPtr(m, v13+int32(4))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v82 = v69
	goto L4
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[1]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v66
	goto L15
L14:
	;
	goto L15
L15:
	;
	v69 = *(*int64)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[2]))
	goto L12
L16:
	;
	return int32(0)
L17:
	;
	v76 = F_GetWALInsertionTimeLineIfSet(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v76 == int32(0) {
		v82 = v72
		goto L4
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v76
	v82 = v72
	goto L4
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v86 == v87 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if base.Ui64(v16) <= base.Ui64(v82) {
		v102 = v83
		v103 = v82
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L3
L24:
	;
	if l4 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v93 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v93)
	v102 = v83
	v103 = v82
	goto L1
L26:
	;
	goto L27
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_read_local_xlog_page_guts[4]))
	if v96 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L16
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_pg_usleep(m, int32(1000))
	mBase = m.M
	goto L2
L31:
	;
	goto L30
L32:
	;
	m.G0 = v13 + int32(48)
	return v120
L33:
	;
	if base.Ui64(v103) < base.Ui64(v16) {
		v120 = int32(-1)
		goto L32
	} else {
		goto L36
	}
L34:
	;
	v112 = int32(_a_F_read_local_xlog_page_guts_2)
	goto L35
L35:
	;
	v114 = v13 + int32(8)
	v115 = F_WALRead(m, l0, l3, l1, v112, v102, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L16
	} else {
		goto L37
	}
L36:
	;
	v112 = base.I32_wrap_i64(v103 - l1)
	goto L35
L37:
	;
	if v115 != 0 {
		v120 = v112
		goto L32
	} else {
		goto L38
	}
L38:
	;
	F_WALReadRaiseError(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L16
	} else {
		goto L39
	}
L39:
	;
	v120 = v112
	goto L32
}
func F_read_local_xlog_page_no_wait(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_read_local_xlog_page_guts(m, l0, l1, l2, l4, int32(0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
