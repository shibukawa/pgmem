package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BufferManagerShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemInit[0]))
	if v2 < v4 {
		v8 = v2
		for {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemInit[1]))
			v13 = v10 + v8*int32(56)
			v14 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v14
			v16 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(-4294967296)
			*(*int64)(unsafe.Add(mBase, uint32(v13))) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v8
			*(*int32)(unsafe.Add(mBase, uint32(v13+int32(36)))) = v16
			v29 = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v29
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemInit[2]))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
			v36 = v32 + v33<<(uint(int32(4))%32)
			v37 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v36))), uint32(v37))
			*(*int64)(unsafe.Add(mBase, uint32(v36)+4)) = v29
			v43 = v8 + int32(1)
			v45 = *(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemInit[0]))
			if v43 < v45 {
				v8 = v43
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	v49 = int32(_a_F_BufferManagerShmemInit_0)
	*(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemInit[3])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemInit[4])) = int32(_a_F_BufferManagerShmemInit_1)
	return
}
func F_BufferUsageAccumDiff(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[0]))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v3 + (v5 - v6)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[1]))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v10 + (v12 - v13)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[2]))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v17 + (v19 - v20)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[3]))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v24 + (v26 - v27)
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[4]))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v31 + (v33 - v34)
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v40 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[5]))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v38 + (v40 - v41)
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v47 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[6]))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v45 + (v47 - v48)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v54 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[7]))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v52 + (v54 - v55)
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v61 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[8]))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v59 + (v61 - v62)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v68 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[9]))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v66 + (v68 - v69)
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	v75 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[10]))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v73 + (v75 - v76)
	v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[11]))
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v80 + (v82 - v83)
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	v89 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[12]))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v87 + (v89 - v90)
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v96 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[13]))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v94 + (v96 - v97)
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v103 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[14]))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v101 + (v103 - v104)
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v110 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[15]))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v108 + (v110 - v111)
	return
}
func F_CheckBufferIsPinnedOnce(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 < v2 {
		v14 = (l0 ^ int32(-1)) << (uint(int32(2)) % 32)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14+v16)))
		if v18 == int32(1) {
			m.G0 = v7 + int32(32)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[0]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v26+v14)))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28
				F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_2), int32(_a_F_CheckBufferIsPinnedOnce_3))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[1]))
		if v39 != int32(-1) {
			v43 = v39 << (uint(int32(4)) % 32)
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_CheckBufferIsPinnedOnce[2])))
			if v46 == l0 {
				v54 = v43 + int32(_a_F_CheckBufferIsPinnedOnce_4)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
				if v55 == int32(1) {
					m.G0 = v7 + int32(32)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[1]))
						if v64 != int32(-1) {
							v68 = v64 << (uint(int32(4)) % 32)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_CheckBufferIsPinnedOnce[2])))
							if v71 == l0 {
								v79 = v68 + int32(_a_F_CheckBufferIsPinnedOnce_4)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
								v82 = v80
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
								F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							} else {
								v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									if v75 == int32(0) {
										v82 = v2
									} else {
										v79 = v75
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
										v82 = v80
									}
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
									F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								if v75 == int32(0) {
									v82 = v2
								} else {
									v79 = v75
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
									v82 = v80
								}
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
								F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
			} else {
				v50 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					if v50 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[1]))
							if v64 != int32(-1) {
								v68 = v64 << (uint(int32(4)) % 32)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_CheckBufferIsPinnedOnce[2])))
								if v71 == l0 {
									v79 = v68 + int32(_a_F_CheckBufferIsPinnedOnce_4)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
									v82 = v80
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
									F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										if v75 == int32(0) {
											v82 = v2
										} else {
											v79 = v75
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
											v82 = v80
										}
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
										F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									if v75 == int32(0) {
										v82 = v2
									} else {
										v79 = v75
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
										v82 = v80
									}
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
									F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
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
						v54 = v50
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
						if v55 == int32(1) {
							m.G0 = v7 + int32(32)
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[1]))
								if v64 != int32(-1) {
									v68 = v64 << (uint(int32(4)) % 32)
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_CheckBufferIsPinnedOnce[2])))
									if v71 == l0 {
										v79 = v68 + int32(_a_F_CheckBufferIsPinnedOnce_4)
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
										v82 = v80
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
										F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											if v75 == int32(0) {
												v82 = v2
											} else {
												v79 = v75
												v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
												v82 = v80
											}
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
											F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										if v75 == int32(0) {
											v82 = v2
										} else {
											v79 = v75
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
											v82 = v80
										}
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
										F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
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
					}
				}
			}
		} else {
			v50 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				if v50 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[1]))
						if v64 != int32(-1) {
							v68 = v64 << (uint(int32(4)) % 32)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_CheckBufferIsPinnedOnce[2])))
							if v71 == l0 {
								v79 = v68 + int32(_a_F_CheckBufferIsPinnedOnce_4)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
								v82 = v80
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
								F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							} else {
								v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									if v75 == int32(0) {
										v82 = v2
									} else {
										v79 = v75
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
										v82 = v80
									}
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
									F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								if v75 == int32(0) {
									v82 = v2
								} else {
									v79 = v75
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
									v82 = v80
								}
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
								F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
					v54 = v50
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
					if v55 == int32(1) {
						m.G0 = v7 + int32(32)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[1]))
							if v64 != int32(-1) {
								v68 = v64 << (uint(int32(4)) % 32)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_CheckBufferIsPinnedOnce[2])))
								if v71 == l0 {
									v79 = v68 + int32(_a_F_CheckBufferIsPinnedOnce_4)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
									v82 = v80
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
									F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										if v75 == int32(0) {
											v82 = v2
										} else {
											v79 = v75
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
											v82 = v80
										}
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
										F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v75 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									if v75 == int32(0) {
										v82 = v2
									} else {
										v79 = v75
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
										v82 = v80
									}
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
									F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+int32(16))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
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
				}
			}
		}
	}
}
func F_MarkBufferDirty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L17
	} else {
		goto L22
	}
L4:
	;
	m.G0 = v8 + int32(16)
	return
L5:
	;
	v12 = int64(0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[0]))
	v19 = v14 + (l0^int32(-1))*int32(56)
	v23 = base.AtomicRmwCmpxchg64(m, v19, int32(24), v12, v12)
	if v23&int64(8388608) == v12 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[1]))
	v39 = int32(56)
	v41 = v38 + l0*v39
	v45 = v41 - int32(32)
	v46 = int64(0)
	v49 = base.AtomicRmwCmpxchg64(m, v45, int32(0), v46, v46)
	v53 = v49
	goto L12
L8:
	;
	goto L4
L9:
	;
	v28 = int32(_a_F_MarkBufferDirty_0)
	v30 = *(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirty[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirty[2])) = v30 + int64(1)
	goto L11
L10:
	;
	goto L11
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v23 | int64(8388608)
	goto L8
L12:
	;
	if v53&int64(4194304) != int64(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v61&int64(8388608) != int64(0) {
		goto L4
	} else {
		goto L20
	}
L14:
	;
	v59 = F_WaitBufHdrUnlocked(m, v41-v39)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v61 = v53
	goto L16
L16:
	;
	v65 = base.AtomicRmwCmpxchg64(m, v45, int32(0), v61, v61|int64(8388608))
	if v61 != v65 {
		v53 = v65
		goto L12
	} else {
		goto L19
	}
L17:
	;
	return
L18:
	;
	v61 = v59
	goto L16
L19:
	;
	goto L13
L20:
	;
	v71 = int32(_a_F_MarkBufferDirty_1)
	v73 = *(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirty[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirty[3])) = v73 + int64(1)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkBufferDirty[4])))
	if v78 != int32(1) {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v81 = int32(_a_F_MarkBufferDirty_2)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[5]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[5])) = v83 + v85
	goto L4
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
	F_errmsg_internal(m, int32(_a_F_MarkBufferDirty_3), v8)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_MarkBufferDirty_4), int32(3177), int32(_a_F_MarkBufferDirty_5))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReadBufferWithoutRelcache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v141 int64
	_ = v141
	var v153 int64
	_ = v153
	var v160 int64
	_ = v160
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v214 int64
	_ = v214
	var v218 int64
	_ = v218
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v20
	v23 = l3 - int32(1)
	v27 = F_smgropen(m, v16+int32(16), int32(-1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		if l2 == int32(-1) {
			v33 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v33
			v35 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v35
			*(*int64)(unsafe.Add(mBase, uint32(v16))) = v35
			*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v33
			if base.Ui32(v23) < base.Ui32(int32(2)) {
				v45 = int32(9)
			} else {
				v45 = int32(1)
			}
			v46 = F_ExtendBufferedRel(m, v16, l1, l4, v45)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v281 = v46
				m.G0 = v16 + int32(112)
				return v281
			}
		} else {
			if base.Ui32(v23) <= base.Ui32(int32(1)) {
				v50 = F_IOContextForStrategy(m, l4)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[0]))
					F_ResourceOwnerEnlarge(m, v54)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_ReservePrivateRefCountEntry(m)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
							*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v59
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v61
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v63
							v68 = v16 + int32(32)
							v69 = F_BufTableHashCode(m, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[1]))
								v79 = v72 + v69&int32(127)<<(uint(int32(7))%32) + int32(_a_F_ReadBufferWithoutRelcache_0)
								v81 = F_LWLockAcquire(m, v79, int32(1))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									v83 = F_BufTableLookup(m, v68, v69)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v83 {
											v88 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[2]))
											v91 = v88 + v83*int32(56)
											v93 = F_PinBuffer(m, v91, l4, int32(0))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												F_LWLockRelease(m, v79)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													if v93 != 0 {
														v193 = v91
														v196 = int32(_a_F_ReadBufferWithoutRelcache_1)
														v198 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[3]))
														v199 = int64(1)
														*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[3])) = v198 + v199
														v202 = int32(1)
														v209 = v50 << (uint(int32(6)) % 32)
														v214 = *(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_c_F_ReadBufferWithoutRelcache[4])))
														*(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_c_F_ReadBufferWithoutRelcache[4]))) = v214 + v199
														v218 = *(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_c_F_ReadBufferWithoutRelcache[5])))
														*(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_c_F_ReadBufferWithoutRelcache[5]))) = v218
														F_pgstat_count_backend_io_op(m, int32(0), v50, int32(2), v202, int64(0))
														mBase = m.M
														*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[6])) = uint8(v202)
														*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[7])) = uint8(v202)
														v230 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[8])))
														if v230 != int32(1) {
															v242 = v193
															v247 = v202
														} else {
															v233 = int32(_a_F_ReadBufferWithoutRelcache_2)
															v235 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[9]))
															v237 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[10]))
															*(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[9])) = v235 + v237
															v242 = v193
															v247 = v202
														}
													} else {
														v242 = v91
														v247 = int32(0)
													}
													v253 = *(*int32)(unsafe.Add(mBase, uint32(v242)+20))
													v255 = v253 + int32(1)
													F_ZeroAndLockBuffer(m, v255, l3, v247)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return int32(0)
													} else {
														v281 = v255
														m.G0 = v16 + int32(112)
														return v281
													}
												}
											}
										} else {
											F_LWLockRelease(m, v79)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return int32(0)
											} else {
												v99 = F_GetVictimBuffer(m, l4, v50)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return int32(0)
												} else {
													v102 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[2]))
													v104 = F_LWLockAcquire(m, v79, int32(0))
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return int32(0)
													} else {
														v106 = int32(56)
														v108 = v102 + v99*v106
														v110 = v108 - v106
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v108-int32(36))))
														v116 = F_BufTableInsert(m, v16+int32(32), v69, v115)
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return int32(0)
														} else {
															if v116 < int32(0) {
																v120 = F_LockBufHdr(m, v110)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return int32(0)
																} else {
																	v122 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v110)+16)) = v122
																	v124 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
																	*(*int64)(unsafe.Add(mBase, uint32(v110)+8)) = v124
																	v126 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
																	*(*int64)(unsafe.Add(mBase, uint32(v110))) = v126
																	v129 = v108 - int32(32)
																	v130 = int64(2181300224)
																	if l5 != 0 {
																		v133 = v130
																	} else {
																		v133 = int64(33816576)
																	}
																	if l1 == int32(3) {
																		v136 = v130
																	} else {
																		v136 = v133
																	}
																	v141 = base.AtomicRmwCmpxchg64(m, v129, int32(0), v120, v136|v120&int64(-38010881))
																	if v120 != v141 {
																		v153 = v141
																		for {
																			v160 = base.AtomicRmwCmpxchg64(m, v129, int32(0), v153, v153&int64(-38010881)|v136)
																			if v153 != v160 {
																				v153 = v160
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	} else {
																	}
																	F_LWLockRelease(m, v79)
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		v242 = v110
																		v247 = int32(0)
																		v253 = *(*int32)(unsafe.Add(mBase, uint32(v242)+20))
																		v255 = v253 + int32(1)
																		F_ZeroAndLockBuffer(m, v255, l3, v247)
																		mBase = m.M
																		v257 = m.ExcPending
																		if v257 != 0 {
																			return int32(0)
																		} else {
																			v281 = v255
																			m.G0 = v16 + int32(112)
																			return v281
																		}
																	}
																}
															} else {
																F_UnpinBuffer(m, v110)
																mBase = m.M
																v179 = m.ExcPending
																if v179 != 0 {
																	return int32(0)
																} else {
																	v180 = int32(0)
																	v182 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[2]))
																	v185 = v182 + v116*int32(56)
																	v187 = F_PinBuffer(m, v185, l4, v180)
																	mBase = m.M
																	v188 = m.ExcPending
																	if v188 != 0 {
																		return int32(0)
																	} else {
																		F_LWLockRelease(m, v79)
																		mBase = m.M
																		v190 = m.ExcPending
																		if v190 != 0 {
																			return int32(0)
																		} else {
																			if v187 == int32(0) {
																				v242 = v185
																				v247 = v180
																			} else {
																				v193 = v185
																				v196 = int32(_a_F_ReadBufferWithoutRelcache_1)
																				v198 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[3]))
																				v199 = int64(1)
																				*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[3])) = v198 + v199
																				v202 = int32(1)
																				v209 = v50 << (uint(int32(6)) % 32)
																				v214 = *(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_c_F_ReadBufferWithoutRelcache[4])))
																				*(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_c_F_ReadBufferWithoutRelcache[4]))) = v214 + v199
																				v218 = *(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_c_F_ReadBufferWithoutRelcache[5])))
																				*(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_c_F_ReadBufferWithoutRelcache[5]))) = v218
																				F_pgstat_count_backend_io_op(m, int32(0), v50, int32(2), v202, int64(0))
																				mBase = m.M
																				*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[6])) = uint8(v202)
																				*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[7])) = uint8(v202)
																				v230 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[8])))
																				if v230 != int32(1) {
																					v242 = v193
																					v247 = v202
																				} else {
																					v233 = int32(_a_F_ReadBufferWithoutRelcache_2)
																					v235 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[9]))
																					v237 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[10]))
																					*(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[9])) = v235 + v237
																					v242 = v193
																					v247 = v202
																				}
																			}
																			v253 = *(*int32)(unsafe.Add(mBase, uint32(v242)+20))
																			v255 = v253 + int32(1)
																			F_ZeroAndLockBuffer(m, v255, l3, v247)
																			mBase = m.M
																			v257 = m.ExcPending
																			if v257 != 0 {
																				return int32(0)
																			} else {
																				v281 = v255
																				m.G0 = v16 + int32(112)
																				return v281
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
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = l1
				if l5 != 0 {
					v262 = int32(112)
				} else {
					v262 = int32(117)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+40)) = uint8(v262)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v27
				v268 = v16 + int32(32)
				if l3 == int32(3) {
					v275 = int32(9)
				} else {
					v275 = int32(8)
				}
				v276 = F_StartReadBuffer(m, v268, v16+int32(28), l2, v275)
				mBase = m.M
				v277 = m.ExcPending
				if v277 != 0 {
					return int32(0)
				} else {
					if v276 != 0 {
						v278 = F_WaitReadBuffers(m, v268)
						mBase = m.M
						v279 = m.ExcPending
						if v279 != 0 {
							return int32(0)
						} else {
							v280 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
							v281 = v280
							m.G0 = v16 + int32(112)
							return v281
						}
					} else {
						v280 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
						v281 = v280
						m.G0 = v16 + int32(112)
						return v281
					}
				}
			}
		}
	}
}
func F_ReleaseBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if l0 != 0 {
		if l0 < int32(0) {
			F_UnpinLocalBuffer(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseBuffer[0]))
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseBuffer[1]))
			v17 = v14 + l0*int32(56)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17-int32(36))))
			F_ResourceOwnerForget(m, v12, base.I64_extend_i32_s(v20+int32(1)), int32(_a_F_ReleaseBuffer_0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_UnpinBufferNoOwner(m, v17-int32(56))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					m.G0 = v5 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(0)
			F_errmsg_internal(m, int32(_a_F_ReleaseBuffer_1), v5)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ReleaseBuffer_2), int32(_a_F_ReleaseBuffer_3), int32(_a_F_ReleaseBuffer_4))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
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
func F_UnlockReleaseBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int64
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var __phi126 int32
	_ = __phi126
	var v127 int32
	_ = v127
	var __phi127 int32
	_ = __phi127
	var v128 int32
	_ = v128
	var __phi128 int32
	_ = __phi128
	var v134 int32
	_ = v134
	var __phi134 int32
	_ = __phi134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v196 int64
	_ = v196
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l0 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L48
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return
L3:
	;
	F_UnpinLocalBuffer(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[0]))
	F_ResourceOwnerForget(m, v21, base.I64_extend_i32_u(l0), int32(_a_F_UnlockReleaseBuffer_0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	goto L2
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[1]))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[2]))
	if v29 != int32(-1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v50 = v27 + l0*int32(56)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
	if v51 == int32(2) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v33 = v29 << (uint(int32(4)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_UnlockReleaseBuffer[3])))
	if v36 == l0 {
		v46 = v29
		v47 = v33 + int32(_a_F_UnlockReleaseBuffer_1)
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v40 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	if v40 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[2]))
	v46 = v45
	v47 = v40
	goto L9
L16:
	;
	v58 = int64(4503599627370496)
	goto L18
L17:
	;
	v58 = int64(17179869184)
	goto L18
L18:
	;
	if v46 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v75 = v50 - int32(56)
	if v51 == int32(3) {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v64 = v46 << (uint(int32(4)) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+uint32(_c_F_UnlockReleaseBuffer[3])))
	if v67 == l0 {
		v73 = v64 + int32(_a_F_UnlockReleaseBuffer_1)
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v71 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v73 = v71
	goto L19
L25:
	;
	v77 = int64(9007199254740992)
	goto L27
L26:
	;
	v77 = v58
	goto L27
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v80 = v78 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v80
	if v80 != 0 {
		v196 = v77
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v200 = base.AtomicRmwSub64(m, v50-int32(32), int32(0), v196)
	v201 = v200 - v196
	F_BufferLockProcessRelease(m, v75, v51, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L43
	}
L29:
	;
	v83 = v77 | int64(1)
	if base.B2i32(base.Ui32(v73) < base.Ui32(int32(_a_F_UnlockReleaseBuffer_1)))|base.B2i32(base.Ui32(int32(_a_F_UnlockReleaseBuffer_2)) <= base.Ui32(v73)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v91 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v91
	v94 = v73 - int32(_a_F_UnlockReleaseBuffer_1)
	*(*int32)(unsafe.Add(mBase, uint32(v94>>(uint(int32(2))%32))+uint32(_c_F_UnlockReleaseBuffer[4]))) = v91
	*(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[5])) = v94 >> (uint(int32(4)) % 32)
	v196 = v83
	goto L28
L31:
	;
	goto L32
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[6]))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	v108 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+8)) = v107 - v108
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v114 = int32(4)
	v118 = v112 & ((v73-v111)>>(uint(v114)%32) + v108)
	v121 = v111 + v118<<(uint(v114)%32)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
	if v122 != v108 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)) = uint8(v178)
	v180 = int32(_a_F_UnlockReleaseBuffer_3)
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[7])) = v182 - int32(1)
	v196 = v83
	goto L28
L34:
	;
	v167 = v73
	goto L33
L35:
	;
	goto L36
L36:
	;
	__phi126 = v73
	__phi127 = v121
	__phi128 = v118
	__phi134 = v112
	v126 = __phi126
	v127 = __phi127
	v128 = __phi128
	v134 = __phi134
	goto L37
L37:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v137 = int32(16)
	v141 = (int32(base.Ui32(v136)>>(uint(v137)%32)) ^ v136) * int32(-2048144789)
	v146 = (int32(base.Ui32(v141)>>(uint(int32(13))%32)) ^ v141) * int32(-1028477387)
	if v128 == (int32(base.Ui32(v146)>>(uint(v137)%32))^v146)&v134 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v167 = v127
	goto L33
L39:
	;
	v167 = v126
	goto L33
L40:
	;
	goto L41
L41:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v127)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v152
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v158 = int32(1)
	v160 = v157 & (v128 + v158)
	v163 = v156 + v160<<(uint(int32(4))%32)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+4)))
	if v164 == v158 {
		__phi126 = v127
		__phi127 = v163
		__phi128 = v160
		__phi134 = v157
		v126 = __phi126
		v127 = __phi127
		v128 = __phi128
		v134 = __phi134
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	if v201&int64(536870912) != int64(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_WakePinCountWaiter(m, v75)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v210 = int32(_a_F_UnlockReleaseBuffer_4)
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[8])) = v212 - int32(1)
	goto L2
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg_internal(m, int32(_a_F_UnlockReleaseBuffer_5), v14)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_UnlockReleaseBuffer_6), int32(_a_F_UnlockReleaseBuffer_7), int32(_a_F_UnlockReleaseBuffer_8))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_buffer_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v6) < base.Ui32(v5) {
		v8 = int32(1)
	} else {
		v8 = int32(-1)
	}
	return v8
}
