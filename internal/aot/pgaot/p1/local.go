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
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v179 int64
	_ = v179
	var v183 int32
	_ = v183
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v204 int64
	_ = v204
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int64
	_ = v223
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
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
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v34&int32(_a_F_FlushLocalBuffer_0) != 0 {
				if l1 == int32(0) {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v43
					v45 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v45
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v43
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[1]))
					v52 = F_smgropen(m, v10+int32(8), v51)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v54 = v52
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)))
						if v56 == int32(0) {
						} else {
							v59 = F_DataChecksumsEnabled(m)
							mBase = m.M
							if v59 == int32(0) {
							} else {
								v62 = F_pg_checksum_page(m, v20, v55)
								mBase = m.M
								*(*uint16)(unsafe.Add(mBase, uint32(v20)+8)) = uint16(v62)
							}
						}
						v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[2])))
						v68 = m.G0
						v70 = v68 - int32(16)
						m.G0 = v70
						if v65 != 0 {
							F___clock_gettime(m, int32(1), v70)
							mBase = m.M
							v74 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+8)))
							v75 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
							v79 = v74 + v75*int64(1000000000)
						} else {
							v79 = int64(0)
						}
						m.G0 = v70 + int32(16)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v20
						F_smgrwritev(m, v54, v84, v83, v10+int32(32), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							v91 = int32(1)
							v95 = int64(8192)
							v99 = m.G0
							v101 = v99 - int32(16)
							m.G0 = v101
							if v79 != int64(0) {
								F___clock_gettime(m, int32(1), v101)
								mBase = m.M
								v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v101)+8)))
								v108 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
								v112 = v107 + (v108*int64(1000000000) - v79)
								v119 = int32(_a_F_FlushLocalBuffer_1)
								v121 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3]))
								v123 = base.I64_div_s(v112, int64(1000))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3])) = v121 + v123
								switch v91 {
								case 0:
									v126 = int32(_a_F_FlushLocalBuffer_2)
									v128 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4])) = v128 + v112
								case 1:
									v131 = int32(_a_F_FlushLocalBuffer_3)
									v133 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5])) = v133 + v112
								default:
								}
								v162 = int32(568)
								v163 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6])) = v163 + v112
								v167 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[7]))
								v174 = int32(0)
								if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v167))|base.B2i32(int32(1)<<(uint(v167)%32)&int32(_a_F_FlushLocalBuffer_4) == v174) == v174 {
									v179 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8])) = v179 + v112
									v183 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v183)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[10])) = uint8(v183)
								} else {
								}
							} else {
							}
							v199 = int32(568)
							v200 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11])) = v200 + base.I64_extend_i32_u(v91)
							v204 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12])) = v204 + v95
							F_pgstat_count_backend_io_op(m, v91, int32(3), int32(7), v91, v95)
							mBase = m.M
							v209 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v209)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[13])) = uint8(v209)
							m.G0 = v101 + int32(16)
							v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v217 & int32(-142606337)
							v221 = int32(_a_F_FlushLocalBuffer_5)
							v223 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14])) = v223 + int64(1)
							m.G0 = v10 + int32(48)
							return
						}
					}
				} else {
					v54 = l1
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)))
					if v56 == int32(0) {
					} else {
						v59 = F_DataChecksumsEnabled(m)
						mBase = m.M
						if v59 == int32(0) {
						} else {
							v62 = F_pg_checksum_page(m, v20, v55)
							mBase = m.M
							*(*uint16)(unsafe.Add(mBase, uint32(v20)+8)) = uint16(v62)
						}
					}
					v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[2])))
					v68 = m.G0
					v70 = v68 - int32(16)
					m.G0 = v70
					if v65 != 0 {
						F___clock_gettime(m, int32(1), v70)
						mBase = m.M
						v74 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+8)))
						v75 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
						v79 = v74 + v75*int64(1000000000)
					} else {
						v79 = int64(0)
					}
					m.G0 = v70 + int32(16)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v20
					F_smgrwritev(m, v54, v84, v83, v10+int32(32), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v91 = int32(1)
						v95 = int64(8192)
						v99 = m.G0
						v101 = v99 - int32(16)
						m.G0 = v101
						if v79 != int64(0) {
							F___clock_gettime(m, int32(1), v101)
							mBase = m.M
							v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v101)+8)))
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
							v112 = v107 + (v108*int64(1000000000) - v79)
							v119 = int32(_a_F_FlushLocalBuffer_1)
							v121 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3]))
							v123 = base.I64_div_s(v112, int64(1000))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3])) = v121 + v123
							switch v91 {
							case 0:
								v126 = int32(_a_F_FlushLocalBuffer_2)
								v128 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4])) = v128 + v112
							case 1:
								v131 = int32(_a_F_FlushLocalBuffer_3)
								v133 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5])) = v133 + v112
							default:
							}
							v162 = int32(568)
							v163 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6])) = v163 + v112
							v167 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[7]))
							v174 = int32(0)
							if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v167))|base.B2i32(int32(1)<<(uint(v167)%32)&int32(_a_F_FlushLocalBuffer_4) == v174) == v174 {
								v179 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8])) = v179 + v112
								v183 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v183)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[10])) = uint8(v183)
							} else {
							}
						} else {
						}
						v199 = int32(568)
						v200 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11])) = v200 + base.I64_extend_i32_u(v91)
						v204 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12])) = v204 + v95
						F_pgstat_count_backend_io_op(m, v91, int32(3), int32(7), v91, v95)
						mBase = m.M
						v209 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v209)
						*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[13])) = uint8(v209)
						m.G0 = v101 + int32(16)
						v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v217 & int32(-142606337)
						v221 = int32(_a_F_FlushLocalBuffer_5)
						v223 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14])) = v223 + int64(1)
						m.G0 = v10 + int32(48)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v233 = m.ExcPending
				if v233 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_FlushLocalBuffer_6), int32(0))
					mBase = m.M
					v237 = m.ExcPending
					if v237 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_FlushLocalBuffer_7), int32(194), int32(_a_F_FlushLocalBuffer_8))
						mBase = m.M
						v242 = m.ExcPending
						if v242 != 0 {
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
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v34&int32(_a_F_FlushLocalBuffer_0) != 0 {
			if l1 == int32(0) {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v43
				v45 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v45
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v43
				v51 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[1]))
				v52 = F_smgropen(m, v10+int32(8), v51)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					v54 = v52
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)))
					if v56 == int32(0) {
					} else {
						v59 = F_DataChecksumsEnabled(m)
						mBase = m.M
						if v59 == int32(0) {
						} else {
							v62 = F_pg_checksum_page(m, v20, v55)
							mBase = m.M
							*(*uint16)(unsafe.Add(mBase, uint32(v20)+8)) = uint16(v62)
						}
					}
					v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[2])))
					v68 = m.G0
					v70 = v68 - int32(16)
					m.G0 = v70
					if v65 != 0 {
						F___clock_gettime(m, int32(1), v70)
						mBase = m.M
						v74 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+8)))
						v75 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
						v79 = v74 + v75*int64(1000000000)
					} else {
						v79 = int64(0)
					}
					m.G0 = v70 + int32(16)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v20
					F_smgrwritev(m, v54, v84, v83, v10+int32(32), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v91 = int32(1)
						v95 = int64(8192)
						v99 = m.G0
						v101 = v99 - int32(16)
						m.G0 = v101
						if v79 != int64(0) {
							F___clock_gettime(m, int32(1), v101)
							mBase = m.M
							v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v101)+8)))
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
							v112 = v107 + (v108*int64(1000000000) - v79)
							v119 = int32(_a_F_FlushLocalBuffer_1)
							v121 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3]))
							v123 = base.I64_div_s(v112, int64(1000))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3])) = v121 + v123
							switch v91 {
							case 0:
								v126 = int32(_a_F_FlushLocalBuffer_2)
								v128 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4])) = v128 + v112
							case 1:
								v131 = int32(_a_F_FlushLocalBuffer_3)
								v133 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5])) = v133 + v112
							default:
							}
							v162 = int32(568)
							v163 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6])) = v163 + v112
							v167 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[7]))
							v174 = int32(0)
							if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v167))|base.B2i32(int32(1)<<(uint(v167)%32)&int32(_a_F_FlushLocalBuffer_4) == v174) == v174 {
								v179 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8])) = v179 + v112
								v183 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v183)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[10])) = uint8(v183)
							} else {
							}
						} else {
						}
						v199 = int32(568)
						v200 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11])) = v200 + base.I64_extend_i32_u(v91)
						v204 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12])) = v204 + v95
						F_pgstat_count_backend_io_op(m, v91, int32(3), int32(7), v91, v95)
						mBase = m.M
						v209 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v209)
						*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[13])) = uint8(v209)
						m.G0 = v101 + int32(16)
						v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v217 & int32(-142606337)
						v221 = int32(_a_F_FlushLocalBuffer_5)
						v223 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14])) = v223 + int64(1)
						m.G0 = v10 + int32(48)
						return
					}
				}
			} else {
				v54 = l1
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)))
				if v56 == int32(0) {
				} else {
					v59 = F_DataChecksumsEnabled(m)
					mBase = m.M
					if v59 == int32(0) {
					} else {
						v62 = F_pg_checksum_page(m, v20, v55)
						mBase = m.M
						*(*uint16)(unsafe.Add(mBase, uint32(v20)+8)) = uint16(v62)
					}
				}
				v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[2])))
				v68 = m.G0
				v70 = v68 - int32(16)
				m.G0 = v70
				if v65 != 0 {
					F___clock_gettime(m, int32(1), v70)
					mBase = m.M
					v74 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+8)))
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
					v79 = v74 + v75*int64(1000000000)
				} else {
					v79 = int64(0)
				}
				m.G0 = v70 + int32(16)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v20
				F_smgrwritev(m, v54, v84, v83, v10+int32(32), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					v91 = int32(1)
					v95 = int64(8192)
					v99 = m.G0
					v101 = v99 - int32(16)
					m.G0 = v101
					if v79 != int64(0) {
						F___clock_gettime(m, int32(1), v101)
						mBase = m.M
						v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v101)+8)))
						v108 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
						v112 = v107 + (v108*int64(1000000000) - v79)
						v119 = int32(_a_F_FlushLocalBuffer_1)
						v121 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3]))
						v123 = base.I64_div_s(v112, int64(1000))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[3])) = v121 + v123
						switch v91 {
						case 0:
							v126 = int32(_a_F_FlushLocalBuffer_2)
							v128 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[4])) = v128 + v112
						case 1:
							v131 = int32(_a_F_FlushLocalBuffer_3)
							v133 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[5])) = v133 + v112
						default:
						}
						v162 = int32(568)
						v163 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6]))
						*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[6])) = v163 + v112
						v167 = *(*int32)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[7]))
						v174 = int32(0)
						if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v167))|base.B2i32(int32(1)<<(uint(v167)%32)&int32(_a_F_FlushLocalBuffer_4) == v174) == v174 {
							v179 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[8])) = v179 + v112
							v183 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v183)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[10])) = uint8(v183)
						} else {
						}
					} else {
					}
					v199 = int32(568)
					v200 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11]))
					*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[11])) = v200 + base.I64_extend_i32_u(v91)
					v204 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12]))
					*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[12])) = v204 + v95
					F_pgstat_count_backend_io_op(m, v91, int32(3), int32(7), v91, v95)
					mBase = m.M
					v209 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[9])) = uint8(v209)
					*(*uint8)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[13])) = uint8(v209)
					m.G0 = v101 + int32(16)
					v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v217 & int32(-142606337)
					v221 = int32(_a_F_FlushLocalBuffer_5)
					v223 = *(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14]))
					*(*int64)(unsafe.Add(mBase, _c_F_FlushLocalBuffer[14])) = v223 + int64(1)
					m.G0 = v10 + int32(48)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v233 = m.ExcPending
			if v233 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_FlushLocalBuffer_6), int32(0))
				mBase = m.M
				v237 = m.ExcPending
				if v237 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_FlushLocalBuffer_7), int32(194), int32(_a_F_FlushLocalBuffer_8))
					mBase = m.M
					v242 = m.ExcPending
					if v242 != 0 {
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
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int64
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v780 int64
	_ = v780
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
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
	v845 = m.ExcPending
	if v845 != 0 {
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
	v220 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(int32(85)) <= base.Ui32(v220) {
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
	v210 = m.ExcPending
	if v210 != 0 {
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
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v111 = v104 + (v105-int32(1))&v25<<(uint(int32(3))%32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v113 = int32(0)
	if base.B2i32(v112 == v113)|base.B2i32(v112 == v111) == v113 {
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
	v83 = v40
	goto L26
L26:
	;
	v92 = v44 + int32(1)
	if v92 < v83 {
		v40 = v83
		v44 = v92
		goto L22
	} else {
		goto L35
	}
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	if int32(0) < v70 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v83 = v80
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
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+28)) = uint8(v73)
	goto L29
L31:
	;
	goto L32
L32:
	;
	F_CatCacheRemoveCList(m, v34, v60-int32(8))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
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
	v121 = v112
	goto L39
L37:
	;
	goto L38
L38:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[2]))
	if v163 != 0 {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v121-int32(20))))
	if v25 != v132 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	if v129 != v111 {
		v121 = v129
		goto L39
	} else {
		goto L50
	}
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v134 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	F_CatCacheRemoveCTup(m, v34, v121-int32(24))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L10
	} else {
		goto L49
	}
L44:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	if v137 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v144 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v144)
	goto L41
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+32))
	if v140 <= int32(0) {
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
	v165 = v163
	goto L54
L52:
	;
	goto L53
L53:
	;
	goto L18
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v174 != v34 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	if v183 != 0 {
		v165 = v183
		goto L54
	} else {
		goto L62
	}
L57:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
	if v176 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v179 != v25 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v181 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)) = uint8(v181)
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
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_LocalExecuteInvalidationMessage_1), int32(701), int32(_a_F_LocalExecuteInvalidationMessage_2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
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
	v225 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220<<(uint(int32(1))%32))+uint32(_c_F_LocalExecuteInvalidationMessage[3]))))
	if v225 <= int32(0) {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v230 = v225
	goto L68
L68:
	;
	v242 = v230 & int32(_a_F_LocalExecuteInvalidationMessage_3) * int32(12)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242)+uint32(_c_F_LocalExecuteInvalidationMessage[4])))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v242)+uint32(_c_F_LocalExecuteInvalidationMessage[5])))
	m.T0[v248].(func(*base.Module, int32, int32, int32))(m, v245, v220, v228)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L10
	} else {
		goto L70
	}
L69:
	;
	goto L2
L70:
	;
	v253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+uint32(_c_F_LocalExecuteInvalidationMessage[6]))))
	if int32(0) < v253 {
		v230 = v253
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v769
	v771 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v771
	v773 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v774 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v773 | v774<<(uint(int32(16))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v771
	v780 = *(*int64)(unsafe.Add(mBase, uint32(v13)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v780
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[7]))
	if v785 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L73:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L10
	} else {
		goto L239
	}
L74:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v725 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v723 != v725 {
		goto L2
	} else {
		goto L233
	}
L75:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v712 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L76:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v700 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L77:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v375 != 0 {
		goto L108
	} else {
		goto L109
	}
L78:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v260 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v260 != v262 {
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
	v265 = m.ExcPending
	if v265 != 0 {
		goto L10
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[8]))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v269 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v271 = v269
	goto L87
L85:
	;
	goto L86
L86:
	;
	goto L2
L87:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v271-int32(12))))
	if v266 == v282 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v285 = v271 - int32(100)
	F_ResetCatalogCache(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L10
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v364 != 0 {
		v271 = v364
		goto L87
	} else {
		goto L107
	}
L92:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v289 = m.G0
	v291 = v289 - int32(16)
	m.G0 = v291
	if base.Ui32(v288) < base.Ui32(int32(85)) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L91
L94:
	;
	v297 = int32(*(*int16)(unsafe.Add(mBase, uint32(v288<<(uint(int32(1))%32))+uint32(_c_F_LocalExecuteInvalidationMessage[3]))))
	if int32(0) < v297 {
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
	v344 = m.ExcPending
	if v344 != 0 {
		goto L10
	} else {
		goto L104
	}
L97:
	;
	v300 = v297
	goto L100
L98:
	;
	goto L99
L99:
	;
	m.G0 = v291 + int32(16)
	goto L93
L100:
	;
	v313 = v300 & int32(_a_F_LocalExecuteInvalidationMessage_3) * int32(12)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_LocalExecuteInvalidationMessage[4])))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_LocalExecuteInvalidationMessage[5])))
	m.T0[v320].(func(*base.Module, int32, int32, int32))(m, v316, v288, int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L10
	} else {
		goto L102
	}
L101:
	;
	goto L99
L102:
	;
	v325 = int32(*(*int16)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_LocalExecuteInvalidationMessage[6]))))
	if int32(0) < v325 {
		v300 = v325
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v288
	F_errmsg_internal(m, int32(_a_F_LocalExecuteInvalidationMessage_0), v291)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_LocalExecuteInvalidationMessage_4), int32(1903), int32(_a_F_LocalExecuteInvalidationMessage_5))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
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
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v375 != v377 {
		goto L2
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v379 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L110
L112:
	;
	v673 = int32(0)
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[9]))
	if v675 <= v673 {
		goto L2
	} else {
		goto L216
	}
L113:
	;
	F_RelationCacheInvalidate(m)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L10
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v384 = m.G0
	v386 = v384 - int32(16)
	m.G0 = v386
	*(*int32)(unsafe.Add(mBase, uint32(v386)+12)) = v379
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[10]))
	v393 = int32(0)
	v395 = F_hash_search(m, v390, v386+int32(12), v393, v393)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L10
	} else {
		goto L119
	}
L116:
	;
	goto L112
L117:
	;
	m.G0 = v386 + int32(16)
	goto L112
L118:
	;
	v460 = int32(_a_F_LocalExecuteInvalidationMessage_6)
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[11])) = v462 + int32(1)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v397)+32))
	if v466 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L119:
	;
	if v395 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v397 != 0 {
		goto L118
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[12]))
	if v400 <= int32(0) {
		goto L117
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v403 = int32(0)
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[13]))
	if v400 != int32(1) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v413 = v403
	v418 = v2
	goto L128
L126:
	;
	v443 = v403
	goto L127
L127:
	;
	v454 = v405 + v443<<(uint(int32(3))%32)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	if v455 != v456 {
		goto L117
	} else {
		goto L138
	}
L128:
	;
	v424 = v405 + v413<<(uint(int32(3))%32)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	if v425 == v426 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v400&int32(1) == int32(0) {
		goto L117
	} else {
		goto L137
	}
L130:
	;
	v428 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+4)) = uint8(v428)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v424)+8))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	if v430 == v431 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v433 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+12)) = uint8(v433)
	goto L135
L134:
	;
	goto L135
L135:
	;
	v435 = int32(2)
	v436 = v413 + v435
	v438 = v418 + v435
	if v438 != v400&int32(2147483646) {
		v413 = v436
		v418 = v438
		goto L128
	} else {
		goto L136
	}
L136:
	;
	goto L129
L137:
	;
	v443 = v436
	goto L127
L138:
	;
	v458 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v454)+4)) = uint8(v458)
	goto L117
L139:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	if v550 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L140:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v397)+40))
	if v469 == int32(0) {
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[14]))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+20))
	goto L145
L143:
	;
	goto L142
L144:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	if v510 != 0 {
		goto L156
	} else {
		goto L157
	}
L145:
	;
	if base.B2i32(v474 == int32(2)) == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v397)+44))
	if v479 != 0 {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[15]))
	F_ResourceOwnerEnlarge(m, v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v397)+16)) = v484 + int32(1)
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[16]))
	if v489 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[15]))
	F_ResourceOwnerRemember(m, v491, v397, int32(_a_F_LocalExecuteInvalidationMessage_7))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L10
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	F_RelationRebuildRelation(m, v397)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L10
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v397)+16)) = v497 - int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[16]))
	if v502 == int32(0) {
		goto L117
	} else {
		goto L154
	}
L154:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[15]))
	F_ResourceOwnerForget(m, v506, v397, int32(_a_F_LocalExecuteInvalidationMessage_7))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	goto L117
L156:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v510)+72))
	v515 = v513 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v510)+72)) = v515
	if v515 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	goto L158
L158:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v397)+256))
	if v543 != 0 {
		goto L168
	} else {
		goto L169
	}
L159:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	F_smgrclose(m, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L10
	} else {
		goto L167
	}
L160:
	;
	v520 = v510 + int32(76)
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17]))
	if v522 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v510)+76)) = v529
	v531 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, uint32(v510)+80)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v529)+4)) = v520
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18])) = v520
	goto L162
L164:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18]))
	v529 = v524
	goto L163
L165:
	;
	goto L166
L166:
	;
	v526 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17])) = v526
	v529 = v526
	goto L163
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+12)) = int32(0)
	goto L158
L168:
	;
	F_pfree(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L10
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v546 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+26)) = uint8(v546)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+256)) = v546
	goto L117
L171:
	;
	goto L170
L172:
	;
	F_RelationClearRelation(m, v397)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L10
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[14]))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+20))
	goto L176
L175:
	;
	goto L117
L176:
	;
	if base.B2i32(v557 == int32(2)) == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	if v562 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	goto L179
L179:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+25)))
	if v602 != int32(1) {
		goto L196
	} else {
		goto L197
	}
L180:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v562)+72))
	v567 = v565 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+72)) = v567
	if v567 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	goto L182
L182:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v397)+256))
	if v595 != 0 {
		goto L192
	} else {
		goto L193
	}
L183:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	F_smgrclose(m, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L10
	} else {
		goto L191
	}
L184:
	;
	v572 = v562 + int32(76)
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17]))
	if v574 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v562)+76)) = v581
	v583 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, uint32(v562)+80)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(v581)+4)) = v572
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18])) = v572
	goto L186
L188:
	;
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18]))
	v581 = v576
	goto L187
L189:
	;
	goto L190
L190:
	;
	v578 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17])) = v578
	v581 = v578
	goto L187
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+12)) = int32(0)
	goto L182
L192:
	;
	F_pfree(m, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L10
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v598 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+26)) = uint8(v598)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+256)) = v598
	goto L117
L195:
	;
	goto L194
L196:
	;
	F_RelationRebuildRelation(m, v397)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L10
	} else {
		goto L215
	}
L197:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	if v605 != int32(1) {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	if v608 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v608)+72))
	v613 = v611 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v608)+72)) = v613
	if v613 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	goto L201
L201:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v397)+256))
	if v641 != 0 {
		goto L211
	} else {
		goto L212
	}
L202:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	F_smgrclose(m, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L10
	} else {
		goto L210
	}
L203:
	;
	v618 = v608 + int32(76)
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17]))
	if v620 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v608)+76)) = v627
	v629 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, uint32(v608)+80)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v627)+4)) = v618
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18])) = v618
	goto L205
L207:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[18]))
	v627 = v622
	goto L206
L208:
	;
	goto L209
L209:
	;
	v624 = int32(_a_F_LocalExecuteInvalidationMessage_8)
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[17])) = v624
	v627 = v624
	goto L206
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+12)) = int32(0)
	goto L201
L211:
	;
	F_pfree(m, v641)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L10
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v644 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+26)) = uint8(v644)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+256)) = v644
	goto L117
L214:
	;
	goto L213
L215:
	;
	goto L117
L216:
	;
	v679 = v673
	goto L217
L217:
	;
	v689 = v679 << (uint(int32(3)) % 32)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)+uint32(_c_F_LocalExecuteInvalidationMessage[19])))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v689)+uint32(_c_F_LocalExecuteInvalidationMessage[20])))
	m.T0[v692].(func(*base.Module, int32, int32))(m, v690, v691)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L10
	} else {
		goto L219
	}
L218:
	;
	goto L2
L219:
	;
	v696 = v679 + int32(1)
	v698 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[9]))
	if v696 < v698 {
		v679 = v696
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
	v705 = m.ExcPending
	if v705 != 0 {
		goto L10
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v707 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v700 != v707 {
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
	v711 = m.ExcPending
	if v711 != 0 {
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
	v716 = m.ExcPending
	if v716 != 0 {
		goto L10
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[0]))
	if v712 != v718 {
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
	v721 = m.ExcPending
	if v721 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	goto L2
L233:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[21]))
	if v728 <= int32(0) {
		goto L2
	} else {
		goto L234
	}
L234:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v733 = int32(0)
	goto L235
L235:
	;
	v743 = v733 << (uint(int32(3)) % 32)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+uint32(_c_F_LocalExecuteInvalidationMessage[22])))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v743)+uint32(_c_F_LocalExecuteInvalidationMessage[23])))
	m.T0[v745].(func(*base.Module, int32, int32))(m, v744, v731)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L10
	} else {
		goto L237
	}
L236:
	;
	goto L2
L237:
	;
	v749 = v733 + int32(1)
	v751 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[21]))
	if v749 < v751 {
		v733 = v749
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	v757 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v757
	F_errmsg_internal(m, int32(_a_F_LocalExecuteInvalidationMessage_9), v11+int32(-48))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L10
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_LocalExecuteInvalidationMessage_4), int32(901), int32(_a_F_LocalExecuteInvalidationMessage_10))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
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
	v788 = int32(0)
	v790 = F_hash_search(m, v785, v11+int32(-32), v788, v788)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L10
	} else {
		goto L244
	}
L244:
	;
	if v790 == int32(0) {
		goto L242
	} else {
		goto L245
	}
L245:
	;
	v794 = int32(_a_F_LocalExecuteInvalidationMessage_11)
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[24])) = v796 + int32(1)
	F_mdclose(m, v790, int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L10
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v790)+20)) = int32(-1)
	F_mdclose(m, v790, int32(1))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L10
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v790)+24)) = int32(-1)
	F_mdclose(m, v790, int32(2))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L10
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v790)+28)) = int32(-1)
	F_mdclose(m, v790, int32(3))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L10
	} else {
		goto L249
	}
L249:
	;
	v818 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v790)+16)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v790)+32)) = v818
	v822 = int32(_a_F_LocalExecuteInvalidationMessage_11)
	v824 = *(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_LocalExecuteInvalidationMessage[24])) = v824 - int32(1)
	goto L242
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v220
	F_errmsg_internal(m, int32(_a_F_LocalExecuteInvalidationMessage_0), v13)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L10
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_LocalExecuteInvalidationMessage_4), int32(1903), int32(_a_F_LocalExecuteInvalidationMessage_5))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(2)
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
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(2)
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v149 int32
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
	var v164 int32
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
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int64
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	v5 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(32)
	m.G0 = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v32
	v36 = base.I32_wrap_i64(v32) & int32(448)
	v38 = l1 + int32(104)
	goto L1
L1:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(23)))) = uint8(v41)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_complete[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	goto L2
L2:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+23)))
	if v50 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v64 = int32(0)
	v70 = v5
	v71 = v5
	v75 = v5
	v76 = v5
	v77 = v5
	v78 = v5
	v79 = v5
	goto L6
L4:
	;
	v244 = v5
	v245 = v5
	v249 = v5
	v251 = v5
	v252 = v5
	v253 = v5
	v264 = int32(0)
	goto L5
L5:
	;
	if v36 == int32(256) {
		goto L52
	} else {
		goto L53
	}
L6:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_complete[1]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(3))%32)+v64<<(uint(int32(3))%32))))
	v99 = v97 ^ int32(-1)
	v102 = v93 + v99<<(uint(int32(6))%32)
	if v97 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v244 = v222
	v245 = v225
	v249 = v212
	v251 = v216
	v252 = v220
	v253 = v227
	v264 = base.B2i32(v221&int32(255) != int32(0))
	goto L5
L8:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)) = uint8(v120)
	if base.B2i32(v36 == int32(256))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(v32)>>(uint(int64(32))%64))) <= v64) != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_complete[2]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v99<<(uint(int32(2))%32))))
	v118 = v110
	goto L8
L10:
	;
	goto L11
L11:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_local_buffer_readv_complete[3]))
	v118 = v112 + v97<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L12:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	goto L30
L13:
	;
	v188 = int32(1)
	v189 = v120
	v190 = int32(0)
	v191 = int32(134217728)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v128 = F_PageIsVerified(m, v118, v119, l3&int32(4)|int32(2), v30+int32(22))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	if v128 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v162<<(uint(int32(18))%32) | v165 | (v164 | (v160<<(uint(int32(9))%32) | v163)) | v64<<(uint(int32(25))%32)
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v30)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v180
	F_pgaio_result_report(m, v30+int32(8), v38, int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L16
	} else {
		goto L28
	}
L19:
	;
	v133 = int32(2048)
	if l3&int32(1) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v148 = int32(1)
	v149 = int32(16777216)
	if v130&v148 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v137 = int32(0)
	v159 = v137
	v160 = v120
	v161 = int32(134217728)
	v162 = v130
	v163 = v137
	v164 = v133
	v165 = int32(259)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v140 = int32(0)
	base.MemoryFill(m, v118, v140, int32(_a_F_local_buffer_readv_complete_0))
	v145 = int32(1)
	v159 = v145
	v160 = v145
	v161 = int32(16777216)
	v162 = v130
	v163 = v140
	v164 = v133
	v165 = int32(195)
	goto L18
L25:
	;
	v188 = v148
	v189 = v120
	v190 = int32(0)
	v191 = v149
	goto L12
L26:
	;
	goto L27
L27:
	;
	v159 = v148
	v160 = v120
	v161 = v149
	v162 = int32(1)
	v163 = int32(1024)
	v164 = int32(0)
	v165 = int32(195)
	goto L18
L28:
	;
	v188 = v159 | v160
	v189 = v160
	v190 = v128
	v191 = v161
	goto L12
L29:
	;
	if v76&int32(255) != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	F_pgaio_wref_clear(m, v102+int32(36))
	mBase = m.M
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+24)) = v198&int32(-134217729) - int32(1) | v191
	goto L29
L33:
	;
	v211 = v75
	goto L35
L34:
	;
	v211 = v64
	goto L35
L35:
	;
	if v190 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v212 = v211
	goto L38
L37:
	;
	v212 = v75
	goto L38
L38:
	;
	if v70&int32(255) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v215 = v77
	goto L41
L40:
	;
	v215 = v64
	goto L41
L41:
	;
	if v189 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v216 = v215
	goto L44
L43:
	;
	v216 = v77
	goto L44
L44:
	;
	if v71&int32(255) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v219 = v78
	goto L47
L46:
	;
	v219 = v64
	goto L47
L47:
	;
	if v188 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v220 = v78
	goto L50
L49:
	;
	v220 = v219
	goto L50
L50:
	;
	v221 = v190 + v76
	v222 = v189 + v70
	v223 = int32(1)
	v225 = v71 + (v188 ^ v223)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v227 = v226 + v79
	v229 = v64 + v223
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+23)))
	if base.Ui32(v229) < base.Ui32(v230) {
		v64 = v229
		v70 = v222
		v71 = v225
		v75 = v212
		v76 = v221
		v77 = v216
		v78 = v220
		v79 = v227
		goto L6
	} else {
		goto L51
	}
L51:
	;
	goto L7
L52:
	;
	v315 = v253 & int32(255)
	if v315 != 0 {
		goto L74
	} else {
		goto L75
	}
L53:
	;
	v267 = int32(255)
	v268 = v244 & v267
	v270 = v245 & v267
	v271 = int32(0)
	if v268|(base.B2i32(v270 != v271)|v264)&int32(1) == v271 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	if v270 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v285 = int32(259)
	goto L57
L56:
	;
	v285 = int32(195)
	goto L57
L57:
	;
	if v268 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v289 = int32(512)
	goto L60
L59:
	;
	v289 = int32(0)
	goto L60
L60:
	;
	if v264 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v292 = int32(1024)
	goto L63
L62:
	;
	v292 = int32(0)
	goto L63
L63:
	;
	if v270 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v294 = v245
	goto L66
L65:
	;
	v294 = v244
	goto L66
L66:
	;
	if v268 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v301 = v251
	goto L69
L68:
	;
	v301 = v249
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
	v302 = v252
	goto L72
L71:
	;
	v302 = v301
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v253&int32(255)<<(uint(int32(18))%32) | v285 | (v289 | v292 | v294&int32(255)<<(uint(int32(11))%32)) | v302<<(uint(int32(25))%32)
	v307 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v307
	F_pgaio_result_report(m, v30, v38, int32(14))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	goto L52
L74:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	F_pgstat_report_checksum_failures_in_db(m, v316, v315)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L16
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	m.G0 = v30 + int32(32)
	return
L77:
	;
	goto L76
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
