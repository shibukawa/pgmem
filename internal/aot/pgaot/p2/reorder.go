package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReorderBufferAllocChange(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v4 = F_MemoryContextAlloc(m, v2, int32(64))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+48)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+40)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+32)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v8
		return v4
	}
}
func F_ReorderBufferAssignChild(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = F_ReorderBufferTXNByXid(m, l0, l1, v11+int32(10), l3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if v18 == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v30 = F_hash_search(m, v24, v11+int32(12), int32(1), v11+int32(11))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
				if v32 == int32(1) {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v35
					v75 = v36
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
					if v78&int32(2) != 0 {
						m.G0 = v11 + int32(16)
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+192))
						*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v82
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
						*(*int32)(unsafe.Add(mBase, uint32(v82))) = v84
						v87 = v75
						*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
						*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
						*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
						v97 = v87 + int32(188)
						v99 = v15 + int32(160)
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
						if v100 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
							*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
						*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
						*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
						F_ReorderBufferTransferSnapToParent(m, v15, v87)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return
						} else {
							m.G0 = v11 + int32(16)
							return
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v41 = F_MemoryContextAlloc(m, v39, int32(232))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v46 = F__emscripten_memset_bulkmem(m, v41, base.I32_extend8_s(int32(0)), int32(232))
						mBase = m.M
						v48 = v46 + int32(160)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+164)) = v48
						*(*int32)(unsafe.Add(mBase, uint32(v46)+160)) = v48
						v52 = v46 + int32(136)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+140)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v46)+136)) = v52
						v56 = v46 + int32(128)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+132)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v46)+128)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v46)+108)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v46
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v62
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
						*(*int64)(unsafe.Add(mBase, uint32(v64)+16)) = l3
						v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
						*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v66
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v62
						if v68&int32(1) == int32(0) {
							v87 = v64
							*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
							*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
							*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
							v97 = v87 + int32(188)
							v99 = v15 + int32(160)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
							if v100 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
							*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
							F_ReorderBufferTransferSnapToParent(m, v15, v87)
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return
							} else {
								m.G0 = v11 + int32(16)
								return
							}
						} else {
							v75 = v64
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
							if v78&int32(2) != 0 {
								m.G0 = v11 + int32(16)
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+192))
								*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v82
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
								*(*int32)(unsafe.Add(mBase, uint32(v82))) = v84
								v87 = v75
								*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
								*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
								*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
								v97 = v87 + int32(188)
								v99 = v15 + int32(160)
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
								if v100 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
									*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
								*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
								*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
								F_ReorderBufferTransferSnapToParent(m, v15, v87)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		} else {
			if l2 != v18 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v30 = F_hash_search(m, v24, v11+int32(12), int32(1), v11+int32(11))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
					if v32 == int32(1) {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v36
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v35
						v75 = v36
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
						if v78&int32(2) != 0 {
							m.G0 = v11 + int32(16)
							return
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+192))
							*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v82
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
							*(*int32)(unsafe.Add(mBase, uint32(v82))) = v84
							v87 = v75
							*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
							*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
							*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
							v97 = v87 + int32(188)
							v99 = v15 + int32(160)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
							if v100 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
							*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
							*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
							F_ReorderBufferTransferSnapToParent(m, v15, v87)
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return
							} else {
								m.G0 = v11 + int32(16)
								return
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v41 = F_MemoryContextAlloc(m, v39, int32(232))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v46 = F__emscripten_memset_bulkmem(m, v41, base.I32_extend8_s(int32(0)), int32(232))
							mBase = m.M
							v48 = v46 + int32(160)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+164)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v46)+160)) = v48
							v52 = v46 + int32(136)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+140)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v46)+136)) = v52
							v56 = v46 + int32(128)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+132)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v46)+128)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v46)+108)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v46
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v62
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v64)+16)) = l3
							v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v66
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v62
							if v68&int32(1) == int32(0) {
								v87 = v64
								*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
								*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
								*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
								v97 = v87 + int32(188)
								v99 = v15 + int32(160)
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
								if v100 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
									*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
								*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
								*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
								F_ReorderBufferTransferSnapToParent(m, v15, v87)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							} else {
								v75 = v64
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
								if v78&int32(2) != 0 {
									m.G0 = v11 + int32(16)
									return
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+192))
									*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v82
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
									*(*int32)(unsafe.Add(mBase, uint32(v82))) = v84
									v87 = v75
									*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
									*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
									*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
									v97 = v87 + int32(188)
									v99 = v15 + int32(160)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
									if v100 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
										*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
									*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
									*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
									F_ReorderBufferTransferSnapToParent(m, v15, v87)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v22 != 0 {
					v75 = v22
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
					if v78&int32(2) != 0 {
						m.G0 = v11 + int32(16)
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+192))
						*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v82
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
						*(*int32)(unsafe.Add(mBase, uint32(v82))) = v84
						v87 = v75
						*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
						*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
						*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
						v97 = v87 + int32(188)
						v99 = v15 + int32(160)
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
						if v100 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
							*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
						*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
						*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
						F_ReorderBufferTransferSnapToParent(m, v15, v87)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return
						} else {
							m.G0 = v11 + int32(16)
							return
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v30 = F_hash_search(m, v24, v11+int32(12), int32(1), v11+int32(11))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
						if v32 == int32(1) {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v36
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v35
							v75 = v36
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
							if v78&int32(2) != 0 {
								m.G0 = v11 + int32(16)
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+192))
								*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v82
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
								*(*int32)(unsafe.Add(mBase, uint32(v82))) = v84
								v87 = v75
								*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
								*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
								*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
								v97 = v87 + int32(188)
								v99 = v15 + int32(160)
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
								if v100 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
									*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
								*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
								*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
								F_ReorderBufferTransferSnapToParent(m, v15, v87)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							}
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v41 = F_MemoryContextAlloc(m, v39, int32(232))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v46 = F__emscripten_memset_bulkmem(m, v41, base.I32_extend8_s(int32(0)), int32(232))
								mBase = m.M
								v48 = v46 + int32(160)
								*(*int32)(unsafe.Add(mBase, uint32(v46)+164)) = v48
								*(*int32)(unsafe.Add(mBase, uint32(v46)+160)) = v48
								v52 = v46 + int32(136)
								*(*int32)(unsafe.Add(mBase, uint32(v46)+140)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v46)+136)) = v52
								v56 = v46 + int32(128)
								*(*int32)(unsafe.Add(mBase, uint32(v46)+132)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v46)+128)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v46)+108)) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v46
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v62
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
								*(*int64)(unsafe.Add(mBase, uint32(v64)+16)) = l3
								v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
								*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v66
								v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v62
								if v68&int32(1) == int32(0) {
									v87 = v64
									*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
									*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
									*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
									v97 = v87 + int32(188)
									v99 = v15 + int32(160)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
									if v100 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
										*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
									*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
									*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
									F_ReorderBufferTransferSnapToParent(m, v15, v87)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16)
										return
									}
								} else {
									v75 = v64
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
									if v78&int32(2) != 0 {
										m.G0 = v11 + int32(16)
										return
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+192))
										*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v82
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+188))
										*(*int32)(unsafe.Add(mBase, uint32(v82))) = v84
										v87 = v75
										*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
										*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
										*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
										v97 = v87 + int32(188)
										v99 = v15 + int32(160)
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
										if v100 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v99
											*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v99
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v99
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
										*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v106
										*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v97
										*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
										*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v110 + int32(1)
										F_ReorderBufferTransferSnapToParent(m, v15, v87)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16)
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
func F_ReorderBufferChangeSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v5 {
	case 0, 1, 2, 8:
		goto L6
	case 3:
		goto L5
	case 4:
		goto L4
	case 5:
		goto L3
	default:
		v104 = int32(64)
		goto L1
	case 11:
		goto L2
	}
L1:
	;
	return v104
L2:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v104 = v98<<(uint(int32(2))%32) - int32(-64)
	goto L1
L3:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	return (v90+v91)<<(uint(int32(2))%32) + int32(136)
L4:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v83<<(uint(int32(4))%32) - int32(-64)
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v20&int32(3) == int32(0) {
		v44 = v20
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v7 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v12 = v8 + int32(84)
	goto L9
L8:
	;
	v12 = int32(64)
	goto L9
L9:
	;
	if v6 == int32(0) {
		v104 = v12
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	return v12 + v15 + int32(20)
L11:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	return v77 + v78 + int32(73)
L12:
	;
	v77 = v69 - v20
	goto L11
L13:
	;
	v48 = v44
	goto L22
L14:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v28 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v77 = int32(0)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v33 = v20
	goto L18
L18:
	;
	v37 = v33 + int32(1)
	if v37&int32(3) == int32(0) {
		v44 = v37
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v69 = v37
	goto L12
L20:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v42 != 0 {
		v33 = v37
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v57 = int32(-2139062144)
	if (int32(16843008)-v54|v54)&v57 == v57 {
		v48 = v48 + int32(4)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v63 = v48
	goto L25
L24:
	;
	goto L23
L25:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v67 != 0 {
		v63 = v63 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v69 = v63
	goto L12
L27:
	;
	goto L26
}
func F_ReorderBufferCopySnap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+168))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v15 = F_MemoryContextAllocZero(m, v7, (v8+v9)<<(uint(int32(2))%32)+int32(76))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = F__emscripten_memcpy_bulkmem(m, v15, l1, int32(72))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(v20)+44)) = int64(1)
		v24 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v20)+30)) = uint8(v24)
		v28 = v20 + int32(72)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
		v33 = v31 << (uint(int32(2)) % 32)
		if v33 != 0 {
			v34 = F__emscripten_memcpy_bulkmem(m, v28, v30, v33)
			mBase = m.M
			v35 = v34
		} else {
			v35 = v28
		}
		v36 = v35 + v33
		*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v36
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v36))) = v38
		*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(1)
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+164))
		if v42 == int32(0) {
			v74 = int32(1)
		} else {
			v47 = l2 + int32(160)
			if v47 == v42 {
				v74 = int32(1)
			} else {
				v51 = v42
				v54 = v24
				for {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(184))))
					*(*int32)(unsafe.Add(mBase, uint32(v56+v54<<(uint(int32(2))%32)))) = v62
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
					v65 = int32(1)
					v66 = v64 + v65
					*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v66
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
					if v70 != v47 {
						v51 = v70
						v54 = v54 + v65
						continue
					} else {
						break
					}
					break
				}
				v74 = v66
			}
		}
		v78 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
		F_pg_qsort(m, v78, v74, int32(4), int32(185))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l3
			return v20
		}
	}
}
func F_ReorderBufferForget(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l1
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+24)) = l2
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
	if v44 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = F_hash_search(m, v19, v11+int32(12), int32(0), v11+int32(11))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if l1 != v14 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v18 != 0 {
		v41 = v18
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L1
L7:
	;
	return
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	if v27 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v30
	goto L1
L10:
	;
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v34
	if v35 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v41 = v35
	goto L2
L13:
	;
	F_ReorderBufferCleanupTXN(m, l0, v41)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L28
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+172))
	if v47 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+176))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferForget[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v55 = base.B2i32(v53 != int32(0))
	goto L16
L16:
	;
	if v53 != int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_BeginInternalSubTransaction(m, int32(_a_F_ReorderBufferForget_0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v63 = int32(0)
	goto L22
L20:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_LocalExecuteInvalidationMessage(m, v50+v63<<(uint(int32(4))%32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L24
	}
L23:
	;
	if v55 == int32(0) {
		goto L13
	} else {
		goto L26
	}
L24:
	;
	v76 = v63 + int32(1)
	if v76 != v47 {
		v63 = v76
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	F_RollbackAndReleaseCurrentSubTransaction(m)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L13
L28:
	;
	goto L1
}
func F_ReorderBufferIterCompare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v15 int64
	_ = v15
	v7 = l2 + int32(16)
	v8 = int32(40)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v7+l1*v8)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v7+l0*v8)))
	return base.B2i32(base.Ui64(v15) < base.Ui64(v11)) - base.B2i32(base.Ui64(v11) < base.Ui64(v15))
}
func F_ReorderBufferIterTXNFinish(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = int32(0)
	v13 = v7
	goto L4
L2:
	;
	goto L3
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v35 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(32)+v12*int32(40))))
	if v19 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_FileClose(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v25 = v13
	goto L8
L8:
	;
	v27 = v12 + int32(1)
	if base.Ui32(v27) < base.Ui32(v25) {
		v12 = v27
		v13 = v25
		goto L4
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = v24
	goto L8
L11:
	;
	goto L5
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_pfree(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L16
	}
L13:
	;
	if v35 == l1+int32(8) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v44
	F_ReorderBufferFreeChange(m, l0, v35-int32(52), int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	F_pfree(m, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	return
}
func F_ReorderBufferReplay(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int32, l6 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v6 = l5
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = l6
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)) = uint16(v6)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14&int32(16) != 0 {
		F_ReorderBufferStreamTXN(m, l1, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v20&int32(64) != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
				m.T0[v23].(func(*base.Module, int32, int32, int64))(m, l1, l0, v19)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v26 | int32(512)
					F_ReorderBufferTruncateTXN(m, l1, l0, int32(1))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferReplay[0])) = int32(0)
						return
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
				m.T0[v36].(func(*base.Module, int32, int32, int64))(m, l1, l0, v19)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_ReorderBufferCleanupTXN(m, l1, l0)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		if v41 == int32(0) {
			if v14&int32(64) != 0 {
				return
			} else {
				F_ReorderBufferCleanupTXN(m, l1, l0)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v48 = int32(0)
			F_ReorderBufferProcessTXN(m, l1, l0, l2, v41, v48, v48)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ReorderBufferSetBaseSnapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_ReorderBufferTXNByXid(m, l0, l1, v9+int32(10), l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		if v15&int32(2) == int32(0) {
			v44 = v13
			*(*int64)(unsafe.Add(mBase, uint32(v44)+88)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = l3
			v49 = l0 + int32(12)
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v50 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v49
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v49
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v44)+100)) = v49
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = v56
			v59 = v44 + int32(96)
			*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v59
			*(*int32)(unsafe.Add(mBase, uint32(v49))) = v59
			m.G0 = v9 + int32(16)
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v20
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v22 == int32(0) {
				v27 = int32(0)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v34 = F_hash_search(m, v28, v9+int32(12), v27, v9+int32(11))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
					if v36 == int32(1) {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						v40 = v39
					} else {
						v40 = v27
					}
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v41
					v44 = v40
					*(*int64)(unsafe.Add(mBase, uint32(v44)+88)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = l3
					v49 = l0 + int32(12)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v50 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v49
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v44)+100)) = v49
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = v56
					v59 = v44 + int32(96)
					*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v59
					*(*int32)(unsafe.Add(mBase, uint32(v49))) = v59
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				if v20 != v22 {
					v27 = int32(0)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v34 = F_hash_search(m, v28, v9+int32(12), v27, v9+int32(11))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
						if v36 == int32(1) {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
							v40 = v39
						} else {
							v40 = v27
						}
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v41
						v44 = v40
						*(*int64)(unsafe.Add(mBase, uint32(v44)+88)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = l3
						v49 = l0 + int32(12)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v50 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v49
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v44)+100)) = v49
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = v56
						v59 = v44 + int32(96)
						*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v59
						*(*int32)(unsafe.Add(mBase, uint32(v49))) = v59
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v44 = v26
					*(*int64)(unsafe.Add(mBase, uint32(v44)+88)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = l3
					v49 = l0 + int32(12)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v50 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v49
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v44)+100)) = v49
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = v56
					v59 = v44 + int32(96)
					*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v59
					*(*int32)(unsafe.Add(mBase, uint32(v49))) = v59
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	}
}
func F_ReorderBufferXidHasCatalogChanges(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 == int32(0) {
		v17 = int32(0)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = F_hash_search(m, v18, v8+int32(12), v17, v8+int32(11))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
			if v28 == int32(0) {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v31
				v47 = v17
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v35
				if v36 == int32(0) {
					v47 = v17
				} else {
					v42 = v36
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
					v47 = v44 & int32(1)
				}
			}
			m.G0 = v8 + int32(16)
			return v47
		}
	} else {
		if l1 != v11 {
			v17 = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v24 = F_hash_search(m, v18, v8+int32(12), v17, v8+int32(11))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
				if v28 == int32(0) {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v31
					v47 = v17
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v35
					if v36 == int32(0) {
						v47 = v17
					} else {
						v42 = v36
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
						v47 = v44 & int32(1)
					}
				}
				m.G0 = v8 + int32(16)
				return v47
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v15 != 0 {
				v42 = v15
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
				v47 = v44 & int32(1)
			} else {
				v47 = int32(0)
			}
			m.G0 = v8 + int32(16)
			return v47
		}
	}
}
func F_ReorderBufferXidSetCatalogChanges(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	v7 = F_ReorderBufferTXNByXid(m, l0, l1, int32(0), l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v9&int32(1) == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9 | int32(1)
			v18 = v7 + int32(196)
			v20 = l0 + int32(20)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v21 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v20
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+200)) = v20
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+196)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v18
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v33 + int32(1)
		} else {
		}
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
		if v39 == int32(0) {
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			if v42&int32(1) != 0 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v39))) = v42 | int32(1)
				v49 = v39 + int32(196)
				v51 = l0 + int32(20)
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v52 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v51
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v51
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v39)+200)) = v51
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+196)) = v60
				*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v49
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v49
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v64 + int32(1)
			}
		}
		return
	}
}
