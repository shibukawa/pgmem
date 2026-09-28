package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufferIsPermanent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v21 int32
	_ = v21
	if int32(0) <= l0 {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_BufferIsPermanent[0]))
		v11 = int64(0)
		v14 = base.AtomicRmwCmpxchg64(m, v5+l0*int32(56)-int32(32), int32(0), v11, v11)
		v21 = base.I32_wrap_i64(int64(base.Ui64(v14&int64(2147483648)) >> (uint(int64(31)) % 64)))
	} else {
		v21 = int32(0)
	}
	return v21
}
func F_BufferManagerShmemAttach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(_a_F_BufferManagerShmemAttach_0)
	*(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemAttach[0])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemAttach[1])) = int32(_a_F_BufferManagerShmemAttach_1)
	return
}
func F_FlushBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v139 int32
	_ = v139
	var v141 int64
	_ = v141
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v208 int64
	_ = v208
	var v212 int64
	_ = v212
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v16 = F_StartSharedBufferIO(m, l0, v4, int32(1), v4)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(1162)
			v21 = int32(_a_F_FlushBuffer_0)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0])) = v11 + int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v22
			if l1 == int32(0) {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v30
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v32
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v34
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v11)+20))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v34
				v42 = F_smgropen(m, v11+int32(8), int32(-1))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = v42
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1]))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v46+v47<<(uint(int32(13))%32))))
					v52 = int64(0)
					v55 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v52, v52)
					if v55&int64(2147483648) != v52 {
						F_XLogFlush(m, base.I64_rotl(v51, int64(32)))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1]))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v69 = v65 + v66<<(uint(int32(13))%32)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							F_PageSetChecksum(m, v69, v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[2])))
								v77 = m.G0
								v79 = v77 - int32(16)
								m.G0 = v79
								if v74 != 0 {
									F___clock_gettime(m, int32(1), v79)
									mBase = m.M
									v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v79)+8)))
									v84 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
									v88 = v83 + v84*int64(1000000000)
								} else {
									v88 = int64(0)
								}
								m.G0 = v79 + int32(16)
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v69
								F_smgrwritev(m, v44, v93, v92, v11+int32(44), int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v100 = int32(0)
									v102 = int32(1)
									v103 = int64(8192)
									v107 = m.G0
									v109 = v107 - int32(16)
									m.G0 = v109
									if v88 != int64(0) {
										F___clock_gettime(m, int32(1), v109)
										mBase = m.M
										v115 = int64(*(*int32)(unsafe.Add(mBase, uint32(v109)+8)))
										v116 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
										v120 = v115 + (v116*int64(1000000000) - v88)
										v127 = int32(_a_F_FlushBuffer_1)
										v129 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[3]))
										v131 = base.I64_div_s(v120, int64(1000))
										*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[3])) = v129 + v131
										switch v100 {
										case 0:
											v134 = int32(_a_F_FlushBuffer_2)
											v136 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[4]))
											*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[4])) = v136 + v120
										case 1:
											v139 = int32(_a_F_FlushBuffer_3)
											v141 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[5]))
											*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[5])) = v141 + v120
										default:
										}
										v164 = int32(0)
										v166 = l2 << (uint(int32(6)) % 32)
										v171 = *(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[6])))
										*(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[6]))) = v171 + v120
										v175 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[7]))
										if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v175))|base.B2i32(int32(1)<<(uint(v175)%32)&int32(_a_F_FlushBuffer_4) == v164) == v164 {
											v187 = *(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[8])))
											*(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[8]))) = v187 + v120
											v191 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[9])) = uint8(v191)
											*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[10])) = uint8(v191)
										} else {
										}
									} else {
									}
									v203 = l2 << (uint(int32(6)) % 32)
									v208 = *(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[11])))
									*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[11]))) = v208 + base.I64_extend_i32_u(v102)
									v212 = *(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[12])))
									*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[12]))) = v212 + v103
									F_pgstat_count_backend_io_op(m, v100, l2, int32(7), v102, v103)
									mBase = m.M
									v217 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[9])) = uint8(v217)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[13])) = uint8(v217)
									m.G0 = v109 + int32(16)
									v225 = int32(_a_F_FlushBuffer_5)
									v227 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14])) = v227 + int64(1)
									v231 = int32(1)
									F_TerminateBufferIO(m, l0, v231, int64(0), v231, int32(0))
									mBase = m.M
									v236 = m.ExcPending
									if v236 != 0 {
										return
									} else {
										v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
										*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0])) = v238
										m.G0 = v11 + int32(48)
										return
									}
								}
							}
						}
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1]))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v69 = v65 + v66<<(uint(int32(13))%32)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						F_PageSetChecksum(m, v69, v70)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[2])))
							v77 = m.G0
							v79 = v77 - int32(16)
							m.G0 = v79
							if v74 != 0 {
								F___clock_gettime(m, int32(1), v79)
								mBase = m.M
								v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v79)+8)))
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
								v88 = v83 + v84*int64(1000000000)
							} else {
								v88 = int64(0)
							}
							m.G0 = v79 + int32(16)
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v69
							F_smgrwritev(m, v44, v93, v92, v11+int32(44), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v100 = int32(0)
								v102 = int32(1)
								v103 = int64(8192)
								v107 = m.G0
								v109 = v107 - int32(16)
								m.G0 = v109
								if v88 != int64(0) {
									F___clock_gettime(m, int32(1), v109)
									mBase = m.M
									v115 = int64(*(*int32)(unsafe.Add(mBase, uint32(v109)+8)))
									v116 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
									v120 = v115 + (v116*int64(1000000000) - v88)
									v127 = int32(_a_F_FlushBuffer_1)
									v129 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[3]))
									v131 = base.I64_div_s(v120, int64(1000))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[3])) = v129 + v131
									switch v100 {
									case 0:
										v134 = int32(_a_F_FlushBuffer_2)
										v136 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[4]))
										*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[4])) = v136 + v120
									case 1:
										v139 = int32(_a_F_FlushBuffer_3)
										v141 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[5]))
										*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[5])) = v141 + v120
									default:
									}
									v164 = int32(0)
									v166 = l2 << (uint(int32(6)) % 32)
									v171 = *(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[6])))
									*(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[6]))) = v171 + v120
									v175 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[7]))
									if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v175))|base.B2i32(int32(1)<<(uint(v175)%32)&int32(_a_F_FlushBuffer_4) == v164) == v164 {
										v187 = *(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[8])))
										*(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[8]))) = v187 + v120
										v191 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[9])) = uint8(v191)
										*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[10])) = uint8(v191)
									} else {
									}
								} else {
								}
								v203 = l2 << (uint(int32(6)) % 32)
								v208 = *(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[11])))
								*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[11]))) = v208 + base.I64_extend_i32_u(v102)
								v212 = *(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[12])))
								*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[12]))) = v212 + v103
								F_pgstat_count_backend_io_op(m, v100, l2, int32(7), v102, v103)
								mBase = m.M
								v217 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[9])) = uint8(v217)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[13])) = uint8(v217)
								m.G0 = v109 + int32(16)
								v225 = int32(_a_F_FlushBuffer_5)
								v227 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14])) = v227 + int64(1)
								v231 = int32(1)
								F_TerminateBufferIO(m, l0, v231, int64(0), v231, int32(0))
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return
								} else {
									v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0])) = v238
									m.G0 = v11 + int32(48)
									return
								}
							}
						}
					}
				}
			} else {
				v44 = l1
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1]))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v46+v47<<(uint(int32(13))%32))))
				v52 = int64(0)
				v55 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v52, v52)
				if v55&int64(2147483648) != v52 {
					F_XLogFlush(m, base.I64_rotl(v51, int64(32)))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1]))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v69 = v65 + v66<<(uint(int32(13))%32)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						F_PageSetChecksum(m, v69, v70)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[2])))
							v77 = m.G0
							v79 = v77 - int32(16)
							m.G0 = v79
							if v74 != 0 {
								F___clock_gettime(m, int32(1), v79)
								mBase = m.M
								v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v79)+8)))
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
								v88 = v83 + v84*int64(1000000000)
							} else {
								v88 = int64(0)
							}
							m.G0 = v79 + int32(16)
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v69
							F_smgrwritev(m, v44, v93, v92, v11+int32(44), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v100 = int32(0)
								v102 = int32(1)
								v103 = int64(8192)
								v107 = m.G0
								v109 = v107 - int32(16)
								m.G0 = v109
								if v88 != int64(0) {
									F___clock_gettime(m, int32(1), v109)
									mBase = m.M
									v115 = int64(*(*int32)(unsafe.Add(mBase, uint32(v109)+8)))
									v116 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
									v120 = v115 + (v116*int64(1000000000) - v88)
									v127 = int32(_a_F_FlushBuffer_1)
									v129 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[3]))
									v131 = base.I64_div_s(v120, int64(1000))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[3])) = v129 + v131
									switch v100 {
									case 0:
										v134 = int32(_a_F_FlushBuffer_2)
										v136 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[4]))
										*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[4])) = v136 + v120
									case 1:
										v139 = int32(_a_F_FlushBuffer_3)
										v141 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[5]))
										*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[5])) = v141 + v120
									default:
									}
									v164 = int32(0)
									v166 = l2 << (uint(int32(6)) % 32)
									v171 = *(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[6])))
									*(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[6]))) = v171 + v120
									v175 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[7]))
									if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v175))|base.B2i32(int32(1)<<(uint(v175)%32)&int32(_a_F_FlushBuffer_4) == v164) == v164 {
										v187 = *(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[8])))
										*(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[8]))) = v187 + v120
										v191 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[9])) = uint8(v191)
										*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[10])) = uint8(v191)
									} else {
									}
								} else {
								}
								v203 = l2 << (uint(int32(6)) % 32)
								v208 = *(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[11])))
								*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[11]))) = v208 + base.I64_extend_i32_u(v102)
								v212 = *(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[12])))
								*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[12]))) = v212 + v103
								F_pgstat_count_backend_io_op(m, v100, l2, int32(7), v102, v103)
								mBase = m.M
								v217 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[9])) = uint8(v217)
								*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[13])) = uint8(v217)
								m.G0 = v109 + int32(16)
								v225 = int32(_a_F_FlushBuffer_5)
								v227 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14]))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14])) = v227 + int64(1)
								v231 = int32(1)
								F_TerminateBufferIO(m, l0, v231, int64(0), v231, int32(0))
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return
								} else {
									v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0])) = v238
									m.G0 = v11 + int32(48)
									return
								}
							}
						}
					}
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1]))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v69 = v65 + v66<<(uint(int32(13))%32)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					F_PageSetChecksum(m, v69, v70)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[2])))
						v77 = m.G0
						v79 = v77 - int32(16)
						m.G0 = v79
						if v74 != 0 {
							F___clock_gettime(m, int32(1), v79)
							mBase = m.M
							v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v79)+8)))
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
							v88 = v83 + v84*int64(1000000000)
						} else {
							v88 = int64(0)
						}
						m.G0 = v79 + int32(16)
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v69
						F_smgrwritev(m, v44, v93, v92, v11+int32(44), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v100 = int32(0)
							v102 = int32(1)
							v103 = int64(8192)
							v107 = m.G0
							v109 = v107 - int32(16)
							m.G0 = v109
							if v88 != int64(0) {
								F___clock_gettime(m, int32(1), v109)
								mBase = m.M
								v115 = int64(*(*int32)(unsafe.Add(mBase, uint32(v109)+8)))
								v116 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
								v120 = v115 + (v116*int64(1000000000) - v88)
								v127 = int32(_a_F_FlushBuffer_1)
								v129 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[3]))
								v131 = base.I64_div_s(v120, int64(1000))
								*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[3])) = v129 + v131
								switch v100 {
								case 0:
									v134 = int32(_a_F_FlushBuffer_2)
									v136 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[4]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[4])) = v136 + v120
								case 1:
									v139 = int32(_a_F_FlushBuffer_3)
									v141 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[5]))
									*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[5])) = v141 + v120
								default:
								}
								v164 = int32(0)
								v166 = l2 << (uint(int32(6)) % 32)
								v171 = *(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[6])))
								*(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[6]))) = v171 + v120
								v175 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[7]))
								if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v175))|base.B2i32(int32(1)<<(uint(v175)%32)&int32(_a_F_FlushBuffer_4) == v164) == v164 {
									v187 = *(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[8])))
									*(*int64)(unsafe.Add(mBase, uint32(v166)+uint32(_c_F_FlushBuffer[8]))) = v187 + v120
									v191 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[9])) = uint8(v191)
									*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[10])) = uint8(v191)
								} else {
								}
							} else {
							}
							v203 = l2 << (uint(int32(6)) % 32)
							v208 = *(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[11])))
							*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[11]))) = v208 + base.I64_extend_i32_u(v102)
							v212 = *(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[12])))
							*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_FlushBuffer[12]))) = v212 + v103
							F_pgstat_count_backend_io_op(m, v100, l2, int32(7), v102, v103)
							mBase = m.M
							v217 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[9])) = uint8(v217)
							*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[13])) = uint8(v217)
							m.G0 = v109 + int32(16)
							v225 = int32(_a_F_FlushBuffer_5)
							v227 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14]))
							*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14])) = v227 + int64(1)
							v231 = int32(1)
							F_TerminateBufferIO(m, l0, v231, int64(0), v231, int32(0))
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return
							} else {
								v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
								*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0])) = v238
								m.G0 = v11 + int32(48)
								return
							}
						}
					}
				}
			}
		} else {
			m.G0 = v11 + int32(48)
			return
		}
	}
}
func F_LockBufferInternal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if int32(0) <= l0 {
		if base.Ui32(int32(3)) <= base.Ui32(l1-int32(1)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
				F_errmsg_internal(m, int32(_a_F_LockBufferInternal_0), v6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_LockBufferInternal_1), int32(_a_F_LockBufferInternal_2), int32(_a_F_LockBufferInternal_3))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferInternal[0]))
			v16 = int32(56)
			F_BufferLockAcquire(m, l0, v15+l0*v16-v16, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_UnpinBufferNoOwner(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var __phi82 int32
	_ = __phi82
	var v83 int32
	_ = v83
	var __phi83 int32
	_ = __phi83
	var v85 int32
	_ = v85
	var __phi85 int32
	_ = __phi85
	var v86 int32
	_ = v86
	var __phi86 int32
	_ = __phi86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = v7 + int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[0]))
	if v11 != int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v27 = v25 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v27
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v15 = v11 << (uint(int32(4)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_UnpinBufferNoOwner[1])))
	if v18 == v9 {
		v24 = v15 + int32(_a_F_UnpinBufferNoOwner_0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v22 = F_GetPrivateRefCountEntrySlow(m, v9, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	return
L7:
	;
	v24 = v22
	goto L1
L8:
	;
	v33 = base.AtomicRmwSub64(m, l0, int32(24), int64(1))
	if v33&int64(536870912) != int64(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	return
L11:
	;
	F_WakePinCountWaiter(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.B2i32(base.Ui32(v24) < base.Ui32(int32(_a_F_UnpinBufferNoOwner_0)))|base.B2i32(base.Ui32(int32(_a_F_UnpinBufferNoOwner_1)) <= base.Ui32(v24)) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v47
	v50 = v24 - int32(_a_F_UnpinBufferNoOwner_0)
	*(*int32)(unsafe.Add(mBase, uint32(v50>>(uint(int32(2))%32))+uint32(_c_F_UnpinBufferNoOwner[2]))) = v47
	*(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[3])) = v50 >> (uint(int32(4)) % 32)
	return
L16:
	;
	goto L17
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[4]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v64 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v63 - v64
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v70 = int32(4)
	v74 = v68 & ((v24-v67)>>(uint(v70)%32) + v64)
	v77 = v67 + v74<<(uint(v70)%32)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+4)))
	if v78 != v64 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)) = uint8(v124)
	v126 = int32(_a_F_UnpinBufferNoOwner_2)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[5])) = v128 - int32(1)
	goto L10
L19:
	;
	v118 = v24
	goto L18
L20:
	;
	goto L21
L21:
	;
	__phi82 = v24
	__phi83 = v77
	__phi85 = v74
	__phi86 = v68
	v82 = __phi82
	v83 = __phi83
	v85 = __phi85
	v86 = __phi86
	goto L22
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v88 = int32(16)
	v92 = (int32(base.Ui32(v87)>>(uint(v88)%32)) ^ v87) * int32(-2048144789)
	v97 = (int32(base.Ui32(v92)>>(uint(int32(13))%32)) ^ v92) * int32(-1028477387)
	if v85 == (int32(base.Ui32(v97)>>(uint(v88)%32))^v97)&v86 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v118 = v83
	goto L18
L24:
	;
	v118 = v82
	goto L18
L25:
	;
	goto L26
L26:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v103
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	*(*int64)(unsafe.Add(mBase, uint32(v82))) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v109 = int32(1)
	v111 = v108 & (v85 + v109)
	v114 = v107 + v111<<(uint(int32(4))%32)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
	if v115 == v109 {
		__phi82 = v83
		__phi83 = v114
		__phi85 = v111
		__phi86 = v108
		v82 = __phi82
		v83 = __phi83
		v85 = __phi85
		v86 = __phi86
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
}
func F_buffer_readv_report(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	v12 = m.G0
	v14 = v12 - int32(208)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = v14 + int32(136)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_buffer_readv_report[0]))
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
	if v26&int32(256) != 0 {
		v29 = v24
	} else {
		v29 = int32(-1)
	}
	F_GetRelationPath(m, v19, v20, v21, v22, v29, base.I32_extend8_s(v26))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return
	} else {
		v35 = v16 + v17 - int32(1)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v38 = int32(base.Ui32(v36) >> (uint(int32(25)) % 32))
		v40 = int32(base.Ui32(v36) >> (uint(int32(18)) % 32))
		v42 = int32(base.Ui32(v36) >> (uint(int32(11)) % 32))
		v44 = v42 & int32(127)
		v45 = int32(1536)
		v46 = v36 & v45
		if v46 == v45 {
			v50 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				if v50 == int32(0) {
					m.G0 = v14 + int32(208)
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v16
						v60 = int32(127)
						v61 = v40 & v60
						*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v61
						v64 = v42 & v60
						*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v64
						F_errmsg(m, int32(_a_F_buffer_readv_report_0), v14+int32(32))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							if base.Ui32(int32(2)) <= base.Ui32(v44) {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v16 + v38
								v78 = F_errdetail(m, int32(_a_F_buffer_readv_report_1), v14+int32(16))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									v82 = v61 + v64 - int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = v82
									F_errhint_plural(m, int32(_a_F_buffer_readv_report_2), int32(_a_F_buffer_readv_report_3), v82, v14)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										v161 = int32(_a_F_buffer_readv_report_4)
										F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return
										} else {
											m.G0 = v14 + int32(208)
											return
										}
									}
								}
							} else {
								v82 = v61 + v64 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v82
								F_errhint_plural(m, int32(_a_F_buffer_readv_report_2), int32(_a_F_buffer_readv_report_3), v82, v14)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									v161 = int32(_a_F_buffer_readv_report_4)
									F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return
									} else {
										m.G0 = v14 + int32(208)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			if v36&int32(448) == int32(256) {
				v109 = v44
				v110 = int32(_a_F_buffer_readv_report_7)
				v111 = int32(_a_F_buffer_readv_report_8)
				v112 = int32(_a_F_buffer_readv_report_9)
				v113 = int32(_a_F_buffer_readv_report_10)
			} else {
				if v46 == int32(512) {
					v109 = v44
					v110 = int32(_a_F_buffer_readv_report_11)
					v111 = int32(_a_F_buffer_readv_report_12)
					v112 = int32(_a_F_buffer_readv_report_1)
					v113 = int32(_a_F_buffer_readv_report_13)
				} else {
					v109 = v40 & int32(127)
					v110 = int32(_a_F_buffer_readv_report_14)
					v111 = int32(_a_F_buffer_readv_report_15)
					v112 = int32(_a_F_buffer_readv_report_16)
					v113 = int32(_a_F_buffer_readv_report_17)
				}
			}
			v115 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return
			} else {
				if v115 == int32(0) {
					m.G0 = v14 + int32(208)
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return
					} else {
						if v109 == int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v16 + v38
							*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v14 + int32(136)
							F_errmsg_internal(m, v110, v14-int32(-64))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return
							} else {
								v161 = int32(_a_F_buffer_readv_report_18)
								F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return
								} else {
									m.G0 = v14 + int32(208)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v16
							*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v109
							*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v14 + int32(136)
							F_errmsg_internal(m, v113, v14+int32(112))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								v144 = int32(_a_F_buffer_readv_report_18)
								if v109 == int32(0) {
									v161 = v144
									F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return
									} else {
										m.G0 = v14 + int32(208)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v16 + v38
									F_errdetail_internal(m, v112, v14+int32(96))
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v109 - int32(1)
										F_errhint_internal(m, v111, v14+int32(80))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return
										} else {
											v161 = v144
											F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return
											} else {
												m.G0 = v14 + int32(208)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_show_buffer_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int64
	_ = v24
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int64
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v256 int64
	_ = v256
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v294 int64
	_ = v294
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int64
	_ = v320
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int64
	_ = v355
	var v357 int32
	_ = v357
	var v360 int64
	_ = v360
	var v362 int32
	_ = v362
	var v365 int64
	_ = v365
	var v367 int32
	_ = v367
	var v370 int64
	_ = v370
	var v372 int32
	_ = v372
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v380 int64
	_ = v380
	var v382 int32
	_ = v382
	var v385 int64
	_ = v385
	var v387 int32
	_ = v387
	var v390 int64
	_ = v390
	var v392 int32
	_ = v392
	var v395 int64
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v404 int64
	_ = v404
	var v410 int32
	_ = v410
	var v413 int64
	_ = v413
	var v419 int32
	_ = v419
	var v422 int64
	_ = v422
	var v428 int32
	_ = v428
	var v431 int64
	_ = v431
	var v437 int32
	_ = v437
	var v440 int64
	_ = v440
	var v446 int32
	_ = v446
	var v449 int64
	_ = v449
	var v455 int32
	_ = v455
	v11 = m.G0
	v13 = v11 - int32(256)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(256)
	return
L2:
	;
	v19 = int32(1)
	if int64(0) < v15 {
		v34 = v19
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_11), int32(0), v15, l0)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L28
	} else {
		goto L126
	}
L5:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if int64(0) < v35 {
		v47 = v19
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int64(0) < v24 {
		v34 = int32(1)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if int64(0) < v28 {
		v34 = int32(1)
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	v34 = base.B2i32(int64(0) < v31)
	goto L5
L9:
	;
	v48 = int32(1)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if v50 <= int64(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	if int64(0) < v38 {
		v47 = v19
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	if int64(0) < v41 {
		v47 = v19
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	v47 = base.B2i32(int64(0) < v44)
	goto L9
L13:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	v56 = base.B2i32(int64(0) < v53)
	goto L15
L14:
	;
	v56 = v48
	goto L15
L15:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	if v57 == int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	v63 = base.B2i32(v60 != int64(0))
	goto L18
L17:
	;
	v63 = v48
	goto L18
L18:
	;
	v64 = int32(1)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	if v66 == int64(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	v72 = base.B2i32(v69 != int64(0))
	goto L21
L20:
	;
	v72 = v64
	goto L21
L21:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v73 == int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	v79 = base.B2i32(v76 != int64(0))
	goto L24
L23:
	;
	v79 = v64
	goto L24
L24:
	;
	if (v47|v34|v56)&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_ExplainIndentText(m, l0)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v79|(v72|v63) == int32(0) {
		goto L1
	} else {
		goto L85
	}
L28:
	;
	return
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v86, int32(_a_F_show_buffer_usage_0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if v34 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v47 != 0 {
		goto L54
	} else {
		goto L55
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v92, int32(_a_F_show_buffer_usage_10))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if int64(0) < v96 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+240)) = v96
	F_appendStringInfo(m, v99, int32(_a_F_show_buffer_usage_2), v13+int32(240))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int64(0) < v106 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+224)) = v106
	F_appendStringInfo(m, v109, int32(_a_F_show_buffer_usage_3), v13+int32(224))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L28
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if int64(0) < v116 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+208)) = v116
	F_appendStringInfo(m, v119, int32(_a_F_show_buffer_usage_4), v13+int32(208))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L28
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if int64(0) < v126 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+192)) = v126
	F_appendStringInfo(m, v129, int32(_a_F_show_buffer_usage_5), v13+int32(192))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L28
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v56|v47 == int32(0) {
		goto L31
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v139, int32(44))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L28
	} else {
		goto L51
	}
L51:
	;
	goto L31
L52:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v222, int32(10))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L28
	} else {
		goto L84
	}
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v197, int32(_a_F_show_buffer_usage_9))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L28
	} else {
		goto L77
	}
L54:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v144, int32(_a_F_show_buffer_usage_1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L28
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v56 == int32(0) {
		goto L52
	} else {
		goto L76
	}
L57:
	;
	v148 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if int64(0) < v148 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+176)) = v148
	F_appendStringInfo(m, v151, int32(_a_F_show_buffer_usage_2), v13+int32(176))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L28
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	if int64(0) < v158 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+160)) = v158
	F_appendStringInfo(m, v161, int32(_a_F_show_buffer_usage_3), v13+int32(160))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L28
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	if int64(0) < v168 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+144)) = v168
	F_appendStringInfo(m, v171, int32(_a_F_show_buffer_usage_4), v13+int32(144))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L28
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	if int64(0) < v178 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = v178
	F_appendStringInfo(m, v181, int32(_a_F_show_buffer_usage_5), v13+int32(128))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L28
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v56 == int32(0) {
		goto L52
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v190, int32(44))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L28
	} else {
		goto L75
	}
L75:
	;
	goto L53
L76:
	;
	goto L53
L77:
	;
	v201 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if int64(0) < v201 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+112)) = v201
	F_appendStringInfo(m, v204, int32(_a_F_show_buffer_usage_3), v13+int32(112))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L28
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v211 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	if v211 <= int64(0) {
		goto L52
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v211
	F_appendStringInfo(m, v214, int32(_a_F_show_buffer_usage_5), v13+int32(96))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L28
	} else {
		goto L83
	}
L83:
	;
	goto L52
L84:
	;
	goto L27
L85:
	;
	F_ExplainIndentText(m, l0)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L28
	} else {
		goto L86
	}
L86:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v233, int32(_a_F_show_buffer_usage_6))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L28
	} else {
		goto L87
	}
L87:
	;
	if v63 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v72 != 0 {
		goto L103
	} else {
		goto L104
	}
L89:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v239, int32(_a_F_show_buffer_usage_10))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L28
	} else {
		goto L90
	}
L90:
	;
	v243 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	if v243 != int64(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+80)) = base.F64_div(base.F64_convert_i64_s(v243), float64(1e+06))
	F_appendStringInfo(m, v246, int32(_a_F_show_buffer_usage_7), v13+int32(80))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L28
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v256 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	if v256 != int64(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L93
L95:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+64)) = base.F64_div(base.F64_convert_i64_s(v256), float64(1e+06))
	F_appendStringInfo(m, v259, int32(_a_F_show_buffer_usage_8), v13-int32(-64))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L28
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v79|v72 == int32(0) {
		goto L88
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v272, int32(44))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L28
	} else {
		goto L100
	}
L100:
	;
	goto L88
L101:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v345, int32(10))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L28
	} else {
		goto L125
	}
L102:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v316, int32(_a_F_show_buffer_usage_9))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L28
	} else {
		goto L118
	}
L103:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v277, int32(_a_F_show_buffer_usage_1))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L28
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v79 == int32(0) {
		goto L101
	} else {
		goto L117
	}
L106:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	if v281 != int64(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+48)) = base.F64_div(base.F64_convert_i64_s(v281), float64(1e+06))
	F_appendStringInfo(m, v284, int32(_a_F_show_buffer_usage_7), v13+int32(48))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L28
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	if v294 != int64(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = base.F64_div(base.F64_convert_i64_s(v294), float64(1e+06))
	F_appendStringInfo(m, v297, int32(_a_F_show_buffer_usage_8), v13+int32(32))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L28
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v79 == int32(0) {
		goto L101
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v309, int32(44))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L28
	} else {
		goto L116
	}
L116:
	;
	goto L102
L117:
	;
	goto L102
L118:
	;
	v320 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v320 != int64(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = base.F64_div(base.F64_convert_i64_s(v320), float64(1e+06))
	F_appendStringInfo(m, v323, int32(_a_F_show_buffer_usage_7), v13+int32(16))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L28
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v333 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	if v333 == int64(0) {
		goto L101
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = base.F64_div(base.F64_convert_i64_s(v333), float64(1e+06))
	F_appendStringInfo(m, v336, int32(_a_F_show_buffer_usage_8), v13)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L28
	} else {
		goto L124
	}
L124:
	;
	goto L101
L125:
	;
	goto L1
L126:
	;
	v355 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_12), int32(0), v355, l0)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L28
	} else {
		goto L127
	}
L127:
	;
	v360 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_13), int32(0), v360, l0)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L28
	} else {
		goto L128
	}
L128:
	;
	v365 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_14), int32(0), v365, l0)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L28
	} else {
		goto L129
	}
L129:
	;
	v370 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_15), int32(0), v370, l0)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L28
	} else {
		goto L130
	}
L130:
	;
	v375 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_16), int32(0), v375, l0)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L28
	} else {
		goto L131
	}
L131:
	;
	v380 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_17), int32(0), v380, l0)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L28
	} else {
		goto L132
	}
L132:
	;
	v385 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_18), int32(0), v385, l0)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L28
	} else {
		goto L133
	}
L133:
	;
	v390 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_19), int32(0), v390, l0)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L28
	} else {
		goto L134
	}
L134:
	;
	v395 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_20), int32(0), v395, l0)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L28
	} else {
		goto L135
	}
L135:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_show_buffer_usage[0])))
	if v399 != int32(1) {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v404 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_21), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v404), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L28
	} else {
		goto L137
	}
L137:
	;
	v413 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_23), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v413), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L28
	} else {
		goto L138
	}
L138:
	;
	v422 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_24), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v422), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L28
	} else {
		goto L139
	}
L139:
	;
	v431 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_25), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v431), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L28
	} else {
		goto L140
	}
L140:
	;
	v440 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_26), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v440), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L28
	} else {
		goto L141
	}
L141:
	;
	v449 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_27), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v449), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L28
	} else {
		goto L142
	}
L142:
	;
	goto L1
}
