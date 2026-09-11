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
	var v46 int64
	_ = v46
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
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v192 int32
	_ = v192
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int64
	_ = v235
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
			if v34&int32(8388608) != 0 {
				if l1 == int32(0) {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v43
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v46
					v51 = *(*int32)(unsafe.Add(mBase, _consts[126]))
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
						v65 = int32(*(*uint8)(unsafe.Add(mBase, _consts[731])))
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
								v119 = int32(4405008)
								v121 = *(*int64)(unsafe.Add(mBase, _consts[744]))
								v123 = base.I64_div_s(v112, int64(1000))
								*(*int64)(unsafe.Add(mBase, _consts[744])) = v121 + v123
								switch v91 {
								case 0:
									v126 = int32(4323632)
									v128 = *(*int64)(unsafe.Add(mBase, _consts[377]))
									*(*int64)(unsafe.Add(mBase, _consts[377])) = v128 + v112
								case 1:
									v131 = int32(4323648)
									v133 = *(*int64)(unsafe.Add(mBase, _consts[379]))
									*(*int64)(unsafe.Add(mBase, _consts[379])) = v133 + v112
								default:
								}
								v164 = int32(4407552)
								v165 = *(*int64)(unsafe.Add(mBase, _consts[745]))
								*(*int64)(unsafe.Add(mBase, _consts[745])) = v165 + v112
								v169 = *(*int32)(unsafe.Add(mBase, _consts[407]))
								if base.Ui32(int32(16)) < base.Ui32(v169) {
								} else {
									if int32(1)<<(uint(v169)%32)&int32(115186) == int32(0) {
									} else {
										v187 = int32(4404448)
										v188 = *(*int64)(unsafe.Add(mBase, _consts[746]))
										*(*int64)(unsafe.Add(mBase, _consts[746])) = v188 + v112
										v192 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v192)
										*(*uint8)(unsafe.Add(mBase, _consts[738])) = uint8(v192)
									}
								}
							} else {
							}
							v209 = int32(4406592)
							v210 = *(*int64)(unsafe.Add(mBase, _consts[747]))
							*(*int64)(unsafe.Add(mBase, _consts[747])) = v210 + base.I64_extend_i32_u(v91)
							v215 = int32(4405632)
							v216 = *(*int64)(unsafe.Add(mBase, _consts[748]))
							*(*int64)(unsafe.Add(mBase, _consts[748])) = v216 + v95
							F_pgstat_count_backend_io_op(m, v91, int32(3), int32(7), v91, v95)
							mBase = m.M
							v221 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v221)
							*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v221)
							m.G0 = v101 + int32(16)
							v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v229 & int32(-142606337)
							v233 = int32(4323600)
							v235 = *(*int64)(unsafe.Add(mBase, _consts[373]))
							*(*int64)(unsafe.Add(mBase, _consts[373])) = v235 + int64(1)
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
					v65 = int32(*(*uint8)(unsafe.Add(mBase, _consts[731])))
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
							v119 = int32(4405008)
							v121 = *(*int64)(unsafe.Add(mBase, _consts[744]))
							v123 = base.I64_div_s(v112, int64(1000))
							*(*int64)(unsafe.Add(mBase, _consts[744])) = v121 + v123
							switch v91 {
							case 0:
								v126 = int32(4323632)
								v128 = *(*int64)(unsafe.Add(mBase, _consts[377]))
								*(*int64)(unsafe.Add(mBase, _consts[377])) = v128 + v112
							case 1:
								v131 = int32(4323648)
								v133 = *(*int64)(unsafe.Add(mBase, _consts[379]))
								*(*int64)(unsafe.Add(mBase, _consts[379])) = v133 + v112
							default:
							}
							v164 = int32(4407552)
							v165 = *(*int64)(unsafe.Add(mBase, _consts[745]))
							*(*int64)(unsafe.Add(mBase, _consts[745])) = v165 + v112
							v169 = *(*int32)(unsafe.Add(mBase, _consts[407]))
							if base.Ui32(int32(16)) < base.Ui32(v169) {
							} else {
								if int32(1)<<(uint(v169)%32)&int32(115186) == int32(0) {
								} else {
									v187 = int32(4404448)
									v188 = *(*int64)(unsafe.Add(mBase, _consts[746]))
									*(*int64)(unsafe.Add(mBase, _consts[746])) = v188 + v112
									v192 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v192)
									*(*uint8)(unsafe.Add(mBase, _consts[738])) = uint8(v192)
								}
							}
						} else {
						}
						v209 = int32(4406592)
						v210 = *(*int64)(unsafe.Add(mBase, _consts[747]))
						*(*int64)(unsafe.Add(mBase, _consts[747])) = v210 + base.I64_extend_i32_u(v91)
						v215 = int32(4405632)
						v216 = *(*int64)(unsafe.Add(mBase, _consts[748]))
						*(*int64)(unsafe.Add(mBase, _consts[748])) = v216 + v95
						F_pgstat_count_backend_io_op(m, v91, int32(3), int32(7), v91, v95)
						mBase = m.M
						v221 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v221)
						*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v221)
						m.G0 = v101 + int32(16)
						v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v229 & int32(-142606337)
						v233 = int32(4323600)
						v235 = *(*int64)(unsafe.Add(mBase, _consts[373]))
						*(*int64)(unsafe.Add(mBase, _consts[373])) = v235 + int64(1)
						m.G0 = v10 + int32(48)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v245 = m.ExcPending
				if v245 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(212449), int32(0))
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return
					} else {
						F_errfinish(m, int32(467641), int32(194), int32(213224))
						mBase = m.M
						v254 = m.ExcPending
						if v254 != 0 {
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
		if v34&int32(8388608) != 0 {
			if l1 == int32(0) {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v43
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v46
				v51 = *(*int32)(unsafe.Add(mBase, _consts[126]))
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
					v65 = int32(*(*uint8)(unsafe.Add(mBase, _consts[731])))
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
							v119 = int32(4405008)
							v121 = *(*int64)(unsafe.Add(mBase, _consts[744]))
							v123 = base.I64_div_s(v112, int64(1000))
							*(*int64)(unsafe.Add(mBase, _consts[744])) = v121 + v123
							switch v91 {
							case 0:
								v126 = int32(4323632)
								v128 = *(*int64)(unsafe.Add(mBase, _consts[377]))
								*(*int64)(unsafe.Add(mBase, _consts[377])) = v128 + v112
							case 1:
								v131 = int32(4323648)
								v133 = *(*int64)(unsafe.Add(mBase, _consts[379]))
								*(*int64)(unsafe.Add(mBase, _consts[379])) = v133 + v112
							default:
							}
							v164 = int32(4407552)
							v165 = *(*int64)(unsafe.Add(mBase, _consts[745]))
							*(*int64)(unsafe.Add(mBase, _consts[745])) = v165 + v112
							v169 = *(*int32)(unsafe.Add(mBase, _consts[407]))
							if base.Ui32(int32(16)) < base.Ui32(v169) {
							} else {
								if int32(1)<<(uint(v169)%32)&int32(115186) == int32(0) {
								} else {
									v187 = int32(4404448)
									v188 = *(*int64)(unsafe.Add(mBase, _consts[746]))
									*(*int64)(unsafe.Add(mBase, _consts[746])) = v188 + v112
									v192 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v192)
									*(*uint8)(unsafe.Add(mBase, _consts[738])) = uint8(v192)
								}
							}
						} else {
						}
						v209 = int32(4406592)
						v210 = *(*int64)(unsafe.Add(mBase, _consts[747]))
						*(*int64)(unsafe.Add(mBase, _consts[747])) = v210 + base.I64_extend_i32_u(v91)
						v215 = int32(4405632)
						v216 = *(*int64)(unsafe.Add(mBase, _consts[748]))
						*(*int64)(unsafe.Add(mBase, _consts[748])) = v216 + v95
						F_pgstat_count_backend_io_op(m, v91, int32(3), int32(7), v91, v95)
						mBase = m.M
						v221 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v221)
						*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v221)
						m.G0 = v101 + int32(16)
						v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v229 & int32(-142606337)
						v233 = int32(4323600)
						v235 = *(*int64)(unsafe.Add(mBase, _consts[373]))
						*(*int64)(unsafe.Add(mBase, _consts[373])) = v235 + int64(1)
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
				v65 = int32(*(*uint8)(unsafe.Add(mBase, _consts[731])))
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
						v119 = int32(4405008)
						v121 = *(*int64)(unsafe.Add(mBase, _consts[744]))
						v123 = base.I64_div_s(v112, int64(1000))
						*(*int64)(unsafe.Add(mBase, _consts[744])) = v121 + v123
						switch v91 {
						case 0:
							v126 = int32(4323632)
							v128 = *(*int64)(unsafe.Add(mBase, _consts[377]))
							*(*int64)(unsafe.Add(mBase, _consts[377])) = v128 + v112
						case 1:
							v131 = int32(4323648)
							v133 = *(*int64)(unsafe.Add(mBase, _consts[379]))
							*(*int64)(unsafe.Add(mBase, _consts[379])) = v133 + v112
						default:
						}
						v164 = int32(4407552)
						v165 = *(*int64)(unsafe.Add(mBase, _consts[745]))
						*(*int64)(unsafe.Add(mBase, _consts[745])) = v165 + v112
						v169 = *(*int32)(unsafe.Add(mBase, _consts[407]))
						if base.Ui32(int32(16)) < base.Ui32(v169) {
						} else {
							if int32(1)<<(uint(v169)%32)&int32(115186) == int32(0) {
							} else {
								v187 = int32(4404448)
								v188 = *(*int64)(unsafe.Add(mBase, _consts[746]))
								*(*int64)(unsafe.Add(mBase, _consts[746])) = v188 + v112
								v192 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v192)
								*(*uint8)(unsafe.Add(mBase, _consts[738])) = uint8(v192)
							}
						}
					} else {
					}
					v209 = int32(4406592)
					v210 = *(*int64)(unsafe.Add(mBase, _consts[747]))
					*(*int64)(unsafe.Add(mBase, _consts[747])) = v210 + base.I64_extend_i32_u(v91)
					v215 = int32(4405632)
					v216 = *(*int64)(unsafe.Add(mBase, _consts[748]))
					*(*int64)(unsafe.Add(mBase, _consts[748])) = v216 + v95
					F_pgstat_count_backend_io_op(m, v91, int32(3), int32(7), v91, v95)
					mBase = m.M
					v221 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v221)
					*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v221)
					m.G0 = v101 + int32(16)
					v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v229 & int32(-142606337)
					v233 = int32(4323600)
					v235 = *(*int64)(unsafe.Add(mBase, _consts[373]))
					*(*int64)(unsafe.Add(mBase, _consts[373])) = v235 + int64(1)
					m.G0 = v10 + int32(48)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v245 = m.ExcPending
			if v245 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(212449), int32(0))
				mBase = m.M
				v249 = m.ExcPending
				if v249 != 0 {
					return
				} else {
					F_errfinish(m, int32(467641), int32(194), int32(213224))
					mBase = m.M
					v254 = m.ExcPending
					if v254 != 0 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
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
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
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
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int64
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int64
	_ = v791
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
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
	v881 = m.ExcPending
	if v881 != 0 {
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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[158]))
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
	v216 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(int32(85)) <= base.Ui32(v216) {
		goto L1
	} else {
		goto L66
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_consts[1142])))
	if v36 != 0 {
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
	v206 = m.ExcPending
	if v206 != 0 {
		goto L10
	} else {
		goto L63
	}
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	if int32(0) < v37 {
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
	v43 = v37
	v48 = v2
	goto L22
L20:
	;
	goto L21
L21:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v110 = v103 + (v104-int32(1))&v25<<(uint(int32(3))%32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v111 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
	v53 = v50 + v48<<(uint(int32(3))%32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v54 == int32(0) {
		v83 = v43
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v91 = v48 + int32(1)
	if v91 < v83 {
		v43 = v83
		v48 = v91
		goto L22
	} else {
		goto L35
	}
L25:
	;
	if v54 == v53 {
		v83 = v43
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v59 = v54
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	if int32(0) < v69 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v83 = v79
	goto L24
L29:
	;
	if v53 != v68 {
		v59 = v68
		goto L27
	} else {
		goto L34
	}
L30:
	;
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+28)) = uint8(v72)
	goto L29
L31:
	;
	goto L32
L32:
	;
	F_CatCacheRemoveCList(m, v36, v59-int32(8))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
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
	v159 = *(*int32)(unsafe.Add(mBase, _consts[1143]))
	if v159 != 0 {
		goto L51
	} else {
		goto L52
	}
L37:
	;
	if v111 == v110 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v118 = v111
	goto L39
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118-int32(20))))
	if v25 != v128 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L36
L41:
	;
	if v125 != v110 {
		v118 = v125
		goto L39
	} else {
		goto L50
	}
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	if v130 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	F_CatCacheRemoveCTup(m, v36, v118-int32(24))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L49
	}
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+12)) = uint8(v140)
	goto L41
L47:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+32))
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
	v161 = v159
	goto L54
L52:
	;
	goto L53
L53:
	;
	goto L18
L54:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v170 != v36 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	if v179 != 0 {
		v161 = v179
		goto L54
	} else {
		goto L62
	}
L57:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+8)))
	if v172 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v175 != v25 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v177 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+9)) = uint8(v177)
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
	F_errmsg_internal(m, int32(458465), v28)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(468530), int32(701), int32(334832))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
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
	v223 = int32(*(*int16)(unsafe.Add(mBase, uint32(v216<<(uint(int32(1))%32))+uint32(_consts[1144]))))
	if v223 <= int32(0) {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v228 = v223
	goto L68
L68:
	;
	v240 = v228 & int32(65535) * int32(12)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240)+uint32(_consts[1145])))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v240)+uint32(_consts[1146])))
	m.T0[v246].(func(*base.Module, int32, int32, int32))(m, v243, v216, v226)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L10
	} else {
		goto L70
	}
L69:
	;
	goto L2
L70:
	;
	v251 = int32(*(*int16)(unsafe.Add(mBase, uint32(v240)+uint32(_consts[1147]))))
	if int32(0) < v251 {
		v228 = v251
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v780 = v11 + int32(-8)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v780))) = v781
	v783 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v783
	v785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v786 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v785 | v786<<(uint(int32(16))%32)
	v791 = *(*int64)(unsafe.Add(mBase, uint32(v780)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v791
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v783
	v797 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	if v797 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L73:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L10
	} else {
		goto L239
	}
L74:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v731 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v729 != v731 {
		goto L2
	} else {
		goto L233
	}
L75:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v718 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L76:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v706 == int32(0) {
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
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v258 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v258 != v260 {
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
	v263 = m.ExcPending
	if v263 != 0 {
		goto L10
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v266 = *(*int32)(unsafe.Add(mBase, _consts[1141]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	if v267 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v270 = v267
	goto L87
L85:
	;
	goto L86
L86:
	;
	goto L2
L87:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v270-int32(12))))
	if v264 == v280 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v283 = v270 - int32(100)
	F_ResetCatalogCache(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v364 != 0 {
		v270 = v364
		goto L87
	} else {
		goto L107
	}
L92:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v287 = m.G0
	v289 = v287 - int32(16)
	m.G0 = v289
	if base.Ui32(v286) < base.Ui32(int32(85)) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L91
L94:
	;
	v297 = int32(*(*int16)(unsafe.Add(mBase, uint32(v286<<(uint(int32(1))%32))+uint32(_consts[1144]))))
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
	m.G0 = v289 + int32(16)
	goto L93
L100:
	;
	v313 = v300 & int32(65535) * int32(12)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v313)+uint32(_consts[1145])))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v313)+uint32(_consts[1146])))
	m.T0[v320].(func(*base.Module, int32, int32, int32))(m, v316, v286, int32(0))
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
	v325 = int32(*(*int16)(unsafe.Add(mBase, uint32(v313)+uint32(_consts[1147]))))
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
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v286
	F_errmsg_internal(m, int32(458465), v289)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(466931), int32(1903), int32(143404))
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
	v377 = *(*int32)(unsafe.Add(mBase, _consts[158]))
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
	v675 = int32(0)
	v677 = *(*int32)(unsafe.Add(mBase, _consts[1148]))
	if v677 <= v675 {
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
	v390 = *(*int32)(unsafe.Add(mBase, _consts[301]))
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
	v462 = int32(4413276)
	v464 = *(*int32)(unsafe.Add(mBase, _consts[1149]))
	*(*int32)(unsafe.Add(mBase, _consts[1149])) = v464 + int32(1)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v397)+32))
	if v468 == int32(0) {
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
	v400 = *(*int32)(unsafe.Add(mBase, _consts[1150]))
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
	v403 = int32(1)
	v405 = int32(0)
	v407 = *(*int32)(unsafe.Add(mBase, _consts[1151]))
	if v400 != v403 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v413 = v405
	v420 = v2
	goto L128
L126:
	;
	v443 = v405
	goto L127
L127:
	;
	if v400&v403 == int32(0) {
		goto L117
	} else {
		goto L137
	}
L128:
	;
	v424 = v407 + v413<<(uint(int32(3))%32)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	if v425 == v426 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v443 = v438
	goto L127
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
	v431 = v424 + int32(8)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	if v432 == v433 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v435 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+4)) = uint8(v435)
	goto L135
L134:
	;
	goto L135
L135:
	;
	v437 = int32(2)
	v438 = v413 + v437
	v440 = v420 + v437
	if v440 != v400&int32(2147483646) {
		v413 = v438
		v420 = v440
		goto L128
	} else {
		goto L136
	}
L136:
	;
	goto L129
L137:
	;
	v456 = v407 + v443<<(uint(int32(3))%32)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	if v457 != v458 {
		goto L117
	} else {
		goto L138
	}
L138:
	;
	v460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v456)+4)) = uint8(v460)
	goto L117
L139:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	if v552 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L140:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v397)+40))
	if v471 == int32(0) {
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+20))
	goto L145
L143:
	;
	goto L142
L144:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	if v512 != 0 {
		goto L156
	} else {
		goto L157
	}
L145:
	;
	if base.B2i32(v476 == int32(2)) == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v397)+44))
	if v481 != 0 {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerEnlarge(m, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v397)+16)) = v486 + int32(1)
	v491 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v491 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerRemember(m, v493, v397, int32(1690048))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
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
	v498 = m.ExcPending
	if v498 != 0 {
		goto L10
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v397)+16)) = v499 - int32(1)
	v504 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v504 == int32(0) {
		goto L117
	} else {
		goto L154
	}
L154:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerForget(m, v508, v397, int32(1690048))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	goto L117
L156:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v512)+72))
	v517 = v515 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v512)+72)) = v517
	if v517 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	goto L158
L158:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v397)+256))
	if v545 != 0 {
		goto L168
	} else {
		goto L169
	}
L159:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	F_smgrclose(m, v540)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L10
	} else {
		goto L167
	}
L160:
	;
	v522 = v512 + int32(76)
	v524 = *(*int32)(unsafe.Add(mBase, _consts[1152]))
	if v524 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v512)+76)) = v531
	v533 = int32(4348760)
	*(*int32)(unsafe.Add(mBase, uint32(v512)+80)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v531)+4)) = v522
	*(*int32)(unsafe.Add(mBase, _consts[1153])) = v522
	goto L162
L164:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _consts[1153]))
	v531 = v526
	goto L163
L165:
	;
	goto L166
L166:
	;
	v528 = int32(4348760)
	*(*int32)(unsafe.Add(mBase, _consts[1152])) = v528
	v531 = v528
	goto L163
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+12)) = int32(0)
	goto L158
L168:
	;
	F_pfree(m, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L10
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v548 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+26)) = uint8(v548)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+256)) = v548
	goto L117
L171:
	;
	goto L170
L172:
	;
	F_RelationClearRelation(m, v397)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L10
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)+20))
	goto L176
L175:
	;
	goto L117
L176:
	;
	if base.B2i32(v559 == int32(2)) == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	if v564 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	goto L179
L179:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+25)))
	if v604 != int32(1) {
		goto L196
	} else {
		goto L197
	}
L180:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v564)+72))
	v569 = v567 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v564)+72)) = v569
	if v569 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	goto L182
L182:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v397)+256))
	if v597 != 0 {
		goto L192
	} else {
		goto L193
	}
L183:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	F_smgrclose(m, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L10
	} else {
		goto L191
	}
L184:
	;
	v574 = v564 + int32(76)
	v576 = *(*int32)(unsafe.Add(mBase, _consts[1152]))
	if v576 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v564)+76)) = v583
	v585 = int32(4348760)
	*(*int32)(unsafe.Add(mBase, uint32(v564)+80)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v583)+4)) = v574
	*(*int32)(unsafe.Add(mBase, _consts[1153])) = v574
	goto L186
L188:
	;
	v578 = *(*int32)(unsafe.Add(mBase, _consts[1153]))
	v583 = v578
	goto L187
L189:
	;
	goto L190
L190:
	;
	v580 = int32(4348760)
	*(*int32)(unsafe.Add(mBase, _consts[1152])) = v580
	v583 = v580
	goto L187
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+12)) = int32(0)
	goto L182
L192:
	;
	F_pfree(m, v597)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L10
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v600 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+26)) = uint8(v600)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+256)) = v600
	goto L117
L195:
	;
	goto L194
L196:
	;
	F_RelationRebuildRelation(m, v397)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L10
	} else {
		goto L215
	}
L197:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	if v607 != int32(1) {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	if v610 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v610)+72))
	v615 = v613 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v610)+72)) = v615
	if v615 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	goto L201
L201:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v397)+256))
	if v643 != 0 {
		goto L211
	} else {
		goto L212
	}
L202:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	F_smgrclose(m, v638)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L10
	} else {
		goto L210
	}
L203:
	;
	v620 = v610 + int32(76)
	v622 = *(*int32)(unsafe.Add(mBase, _consts[1152]))
	if v622 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v610)+76)) = v629
	v631 = int32(4348760)
	*(*int32)(unsafe.Add(mBase, uint32(v610)+80)) = v631
	*(*int32)(unsafe.Add(mBase, uint32(v629)+4)) = v620
	*(*int32)(unsafe.Add(mBase, _consts[1153])) = v620
	goto L205
L207:
	;
	v624 = *(*int32)(unsafe.Add(mBase, _consts[1153]))
	v629 = v624
	goto L206
L208:
	;
	goto L209
L209:
	;
	v626 = int32(4348760)
	*(*int32)(unsafe.Add(mBase, _consts[1152])) = v626
	v629 = v626
	goto L206
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+12)) = int32(0)
	goto L201
L211:
	;
	F_pfree(m, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L10
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v646 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+26)) = uint8(v646)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+256)) = v646
	goto L117
L214:
	;
	goto L213
L215:
	;
	goto L117
L216:
	;
	v681 = v675
	goto L217
L217:
	;
	v691 = v681 << (uint(int32(3)) % 32)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v691)+uint32(_consts[1154])))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v691)+uint32(_consts[1155])))
	m.T0[v698].(func(*base.Module, int32, int32))(m, v694, v695)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L10
	} else {
		goto L219
	}
L218:
	;
	goto L2
L219:
	;
	v702 = v681 + int32(1)
	v704 = *(*int32)(unsafe.Add(mBase, _consts[1148]))
	if v702 < v704 {
		v681 = v702
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
	v711 = m.ExcPending
	if v711 != 0 {
		goto L10
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v713 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v706 != v713 {
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
	v717 = m.ExcPending
	if v717 != 0 {
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
	v722 = m.ExcPending
	if v722 != 0 {
		goto L10
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v718 != v724 {
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
	v727 = m.ExcPending
	if v727 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	goto L2
L233:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _consts[1156]))
	if v734 <= int32(0) {
		goto L2
	} else {
		goto L234
	}
L234:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v739 = int32(0)
	goto L235
L235:
	;
	v749 = v739 << (uint(int32(3)) % 32)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v749)+uint32(_consts[1157])))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v749)+uint32(_consts[1158])))
	m.T0[v755].(func(*base.Module, int32, int32))(m, v752, v737)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L10
	} else {
		goto L237
	}
L236:
	;
	goto L2
L237:
	;
	v759 = v739 + int32(1)
	v761 = *(*int32)(unsafe.Add(mBase, _consts[1156]))
	if v759 < v761 {
		v739 = v759
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	v767 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v767
	F_errmsg_internal(m, int32(458486), v11+int32(-48))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L10
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(466931), int32(901), int32(381301))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
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
	v800 = int32(0)
	v802 = F_hash_search(m, v797, v11+int32(-32), v800, v800)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L10
	} else {
		goto L244
	}
L244:
	;
	if v802 == int32(0) {
		goto L242
	} else {
		goto L245
	}
L245:
	;
	v806 = int32(4419932)
	v808 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v808 + int32(1)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v802)+36))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v813*int32(80))+uint32(_consts[821])))
	m.T0[v818].(func(*base.Module, int32, int32))(m, v802, int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L10
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+20)) = int32(-1)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v802)+36))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v824*int32(80))+uint32(_consts[821])))
	m.T0[v829].(func(*base.Module, int32, int32))(m, v802, int32(1))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L10
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+24)) = int32(-1)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v802)+36))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v835*int32(80))+uint32(_consts[821])))
	m.T0[v840].(func(*base.Module, int32, int32))(m, v802, int32(2))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L10
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+28)) = int32(-1)
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v802)+36))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v846*int32(80))+uint32(_consts[821])))
	m.T0[v851].(func(*base.Module, int32, int32))(m, v802, int32(3))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L10
	} else {
		goto L249
	}
L249:
	;
	v854 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v802)+16)) = v854
	*(*int32)(unsafe.Add(mBase, uint32(v802)+32)) = v854
	v858 = int32(4419932)
	v860 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v860 - int32(1)
	goto L242
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v216
	F_errmsg_internal(m, int32(458465), v13)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L10
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(466931), int32(1903), int32(143404))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
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
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	v11 = F_palloc(m, int32(48))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = F_pstrdup(m, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = F_pstrdup(m, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v18 = int32(0)
	goto L6
L6:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v18
	if l1&int32(3) == v19 {
		v45 = l1
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v18 = v16
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v78
	*(*float64)(unsafe.Add(mBase, uint32(v11)+40)) = l5
	*(*float64)(unsafe.Add(mBase, uint32(v11)+32)) = l4
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
	v88 = F_palloc(m, int32(8))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L25
	}
L9:
	;
	v78 = v70 - l1
	goto L8
L10:
	;
	v49 = v45
	goto L19
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v29 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v78 = int32(0)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v34 = l1
	goto L15
L15:
	;
	v38 = v34 + int32(1)
	if v38&int32(3) == int32(0) {
		v45 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v70 = v38
	goto L9
L17:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v43 != 0 {
		v34 = v38
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v58 = int32(-2139062144)
	if (int32(16843008)-v55|v55)&v58 == v58 {
		v49 = v49 + int32(4)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v64 = v49
	goto L22
L21:
	;
	goto L20
L22:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v68 != 0 {
		v64 = v64 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v70 = v64
	goto L9
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v11
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = F_lappend(m, v92, v88)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v93
	return
}
func F_local_buffer_readv_complete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v214 int64
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int64
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	v5 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(32)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v30)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v34)
	v39 = v34 & int32(448)
	v41 = l1 + int32(104)
	goto L1
L1:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(23)))) = uint8(v44)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	goto L2
L2:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+23)))
	if v53 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v39 == int32(256) {
		goto L69
	} else {
		goto L70
	}
L4:
	;
	v56 = int32(0)
	v277 = v56
	v281 = v5
	v284 = v5
	v285 = v5
	v286 = v5
	v288 = v5
	v299 = v56
	goto L3
L5:
	;
	goto L6
L6:
	;
	v66 = int32(0)
	v70 = v66
	v71 = v66
	v75 = v5
	v76 = v5
	v78 = v5
	v79 = v5
	v80 = v5
	v82 = v5
	goto L7
L7:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v48+v49<<(uint(int32(3))%32)+v70<<(uint(int32(3))%32))))
	v102 = v100 ^ int32(-1)
	v105 = v96 + v102<<(uint(int32(6))%32)
	if v100 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v277 = v264
	v281 = v246
	v284 = v263
	v285 = v257
	v286 = v261
	v288 = v248
	v299 = base.B2i32(v265&int32(255) != int32(0))
	goto L3
L9:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)) = uint8(v123)
	if base.B2i32(v39 == int32(256))|base.B2i32(v30 <= v70) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v102<<(uint(int32(2))%32))))
	v121 = v113
	goto L9
L11:
	;
	goto L12
L12:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v121 = v115 + v100<<(uint(int32(13))%32) + int32(-8192)
	goto L9
L13:
	;
	if v76&int32(255) != 0 {
		goto L56
	} else {
		goto L57
	}
L14:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	goto L18
L15:
	;
	goto L16
L16:
	;
	v143 = F_PageIsVerified(m, v121, v122, l3&int32(4)|int32(2), v28+int32(22))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v243 = v123
	v245 = int32(0)
	v246 = v75
	v248 = v82
	goto L13
L18:
	;
	F_pgaio_wref_clear(m, v105+int32(36))
	mBase = m.M
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v130&int32(-134217729) - int32(1) | int32(134217728)
	goto L17
L21:
	;
	return
L22:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	if v143 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	if v190 != 0 {
		goto L38
	} else {
		goto L39
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = int64(0)
	v190 = v158
	v191 = v157
	v192 = int32(16777216)
	v193 = int32(195)
	v194 = int32(0)
	v195 = v158 << (uint(int32(11)) % 32)
	goto L23
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = int64(0)
	v180 = int32(0)
	v190 = v180
	v191 = v180
	v192 = int32(134217728)
	v193 = int32(259)
	v194 = int32(1)
	v195 = int32(2048)
	goto L23
L26:
	;
	if v145&int32(1) != 0 {
		goto L24
	} else {
		goto L32
	}
L27:
	;
	if l3&int32(1) == int32(0) {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v157 = v145
	v158 = int32(0)
	goto L26
L30:
	;
	v150 = int32(0)
	v154 = F__emscripten_memset_bulkmem(m, v121, base.I32_extend8_s(v150), int32(8192))
	mBase = m.M
	goto L31
L31:
	;
	v157 = v150
	v158 = int32(1)
	goto L26
L32:
	;
	if v158 != 0 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	goto L35
L34:
	;
	v243 = v123
	v245 = v157
	v246 = v75
	v248 = v82
	goto L13
L35:
	;
	F_pgaio_wref_clear(m, v105+int32(36))
	mBase = m.M
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v164&int32(-134217729) - int32(1) | int32(16777216)
	goto L34
L38:
	;
	v201 = int32(512)
	goto L40
L39:
	;
	v201 = int32(0)
	goto L40
L40:
	;
	if v191&int32(1) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v206 = int32(1024)
	goto L43
L42:
	;
	v206 = int32(0)
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v145<<(uint(int32(18))%32) | v193 | (v195 | (v201 | v206)) | v70<<(uint(int32(25))%32)
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v214
	F_pgaio_result_report(m, v28+int32(8), v41, int32(16))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	goto L46
L45:
	;
	if v194 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	F_pgaio_wref_clear(m, v105+int32(36))
	mBase = m.M
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v223&int32(-134217729) - int32(1) | v192
	goto L45
L49:
	;
	v243 = v190
	v245 = v191
	v246 = v75
	v248 = v82
	goto L13
L50:
	;
	goto L51
L51:
	;
	if v190 != 0 {
		v243 = int32(1)
		v245 = v191
		v246 = v75
		v248 = v82
		goto L13
	} else {
		goto L52
	}
L52:
	;
	if v75&int32(255) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v239 = v82
	goto L55
L54:
	;
	v239 = v70
	goto L55
L55:
	;
	v243 = int32(0)
	v245 = v191
	v246 = v75 + int32(1)
	v248 = v239
	goto L13
L56:
	;
	v254 = v79
	goto L58
L57:
	;
	v254 = v70
	goto L58
L58:
	;
	v256 = v245 & int32(1)
	if v256 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v257 = v254
	goto L61
L60:
	;
	v257 = v79
	goto L61
L61:
	;
	if v71&int32(255) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v260 = v80
	goto L64
L63:
	;
	v260 = v70
	goto L64
L64:
	;
	if v243 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v261 = v260
	goto L67
L66:
	;
	v261 = v80
	goto L67
L67:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v263 = v262 + v78
	v264 = v243 + v71
	v265 = v76 + v256
	v267 = v70 + int32(1)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+23)))
	if base.Ui32(v267) < base.Ui32(v268) {
		v70 = v267
		v71 = v264
		v75 = v246
		v76 = v265
		v78 = v263
		v79 = v257
		v80 = v261
		v82 = v248
		goto L7
	} else {
		goto L68
	}
L68:
	;
	goto L8
L69:
	;
	v353 = v284 & int32(255)
	if v353 != 0 {
		goto L88
	} else {
		goto L89
	}
L70:
	;
	v302 = int32(255)
	v303 = v281 & v302
	v304 = int32(0)
	if (base.B2i32(v303 != v304)|v299)&int32(1)|v277&v302 == v304 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	if v303 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v320 = int32(259)
	goto L74
L73:
	;
	v320 = int32(195)
	goto L74
L74:
	;
	v323 = v277 & int32(255)
	v324 = int32(0)
	if v299 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v330 = int32(1024)
	goto L77
L76:
	;
	v330 = v324
	goto L77
L77:
	;
	if v303 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v332 = v281
	goto L80
L79:
	;
	v332 = v277
	goto L80
L80:
	;
	if v323 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v339 = v286
	goto L83
L82:
	;
	v339 = v285
	goto L83
L83:
	;
	if v303 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v340 = v288
	goto L86
L85:
	;
	v340 = v339
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v284&int32(255)<<(uint(int32(18))%32) | v320 | (base.B2i32(v323 != v324)<<(uint(int32(9))%32) | v330 | v332&int32(255)<<(uint(int32(11))%32)) | v340<<(uint(int32(25))%32)
	v345 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v345
	F_pgaio_result_report(m, v28, v41, int32(14))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L21
	} else {
		goto L87
	}
L87:
	;
	goto L69
L88:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	F_pgstat_report_checksum_failures_in_db(m, v354, v353)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L21
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	m.G0 = v28 + int32(32)
	return
L91:
	;
	goto L90
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
	var v43 int32
	_ = v43
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
	if v17 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+2)) = uint8(v48)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+3)) = uint8(v3)
	return v43 + int32(8)
L4:
	;
	v43 = v9
	goto L3
L5:
	;
	goto L6
L6:
	;
	v25 = v9
	v26 = base.I64_extend_i32_u(v17)
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = F_MemoryContextAlloc(m, v27, int32(24))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v43 = v29
	goto L3
L9:
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
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
}
