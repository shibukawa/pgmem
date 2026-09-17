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
		*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+48)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+40)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+32)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v8
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = v8
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
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
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
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
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
		v19 = int32(0)
		if base.B2i32(v18 == v19)|base.B2i32(l2 != v18) == v19 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v25 != 0 {
				v75 = v25
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
					v97 = v15 + int32(160)
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
					if v98 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v97
						*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v97
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
					*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v104
					v107 = v87 + int32(188)
					*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v107
					*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v107
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
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v33 = F_hash_search(m, v27, v11+int32(12), int32(1), v11+int32(11))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
					if v35 == int32(1) {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v38
						v75 = v39
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
							v97 = v15 + int32(160)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
							if v98 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v97
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v104
							v107 = v87 + int32(188)
							*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v107
							*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v107
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
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v44 = F_MemoryContextAlloc(m, v42, int32(232))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v46 = int32(0)
							base.MemoryFill(m, v44, v46, int32(232))
							v50 = v44 + int32(160)
							*(*int32)(unsafe.Add(mBase, uint32(v44)+164)) = v50
							*(*int32)(unsafe.Add(mBase, uint32(v44)+160)) = v50
							v54 = v44 + int32(136)
							*(*int32)(unsafe.Add(mBase, uint32(v44)+140)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v44)+136)) = v54
							v58 = v44 + int32(128)
							*(*int32)(unsafe.Add(mBase, uint32(v44)+132)) = v58
							*(*int32)(unsafe.Add(mBase, uint32(v44)+128)) = v58
							*(*int32)(unsafe.Add(mBase, uint32(v44)+108)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v44
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v64
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v66)+16)) = l3
							v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
							*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v68
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v66
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v64
							if v70 == v46 {
								v87 = v66
								*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
								*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
								*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
								v97 = v15 + int32(160)
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
								if v98 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v97
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v104
								v107 = v87 + int32(188)
								*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v107
								*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v107
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
								v75 = v66
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
									v97 = v15 + int32(160)
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
									if v98 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v97
										*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v97
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
									*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v104
									v107 = v87 + int32(188)
									*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v107
									*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v107
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
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = F_hash_search(m, v27, v11+int32(12), int32(1), v11+int32(11))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
				if v35 == int32(1) {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v38
					v75 = v39
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
						v97 = v15 + int32(160)
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
						if v98 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v97
							*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v97
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v104
						v107 = v87 + int32(188)
						*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v107
						*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v107
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
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v44 = F_MemoryContextAlloc(m, v42, int32(232))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v46 = int32(0)
						base.MemoryFill(m, v44, v46, int32(232))
						v50 = v44 + int32(160)
						*(*int32)(unsafe.Add(mBase, uint32(v44)+164)) = v50
						*(*int32)(unsafe.Add(mBase, uint32(v44)+160)) = v50
						v54 = v44 + int32(136)
						*(*int32)(unsafe.Add(mBase, uint32(v44)+140)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v44)+136)) = v54
						v58 = v44 + int32(128)
						*(*int32)(unsafe.Add(mBase, uint32(v44)+132)) = v58
						*(*int32)(unsafe.Add(mBase, uint32(v44)+128)) = v58
						*(*int32)(unsafe.Add(mBase, uint32(v44)+108)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v44
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v64
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						*(*int64)(unsafe.Add(mBase, uint32(v66)+16)) = l3
						v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
						*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v68
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v66
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v64
						if v70 == v46 {
							v87 = v66
							*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v15
							*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
							*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92 | int32(2)
							v97 = v15 + int32(160)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
							if v98 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v97
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v104
							v107 = v87 + int32(188)
							*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v107
							*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v107
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
							v75 = v66
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
								v97 = v15 + int32(160)
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
								if v98 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v97
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v87)+192)) = v97
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+188)) = v104
								v107 = v87 + int32(188)
								*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v107
								*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v107
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v5 {
	case 0, 1, 2, 8:
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v7 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v12 = v8 + int32(84)
		} else {
			v12 = int32(64)
		}
		if v6 == int32(0) {
			v48 = v12
			return v48
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			return v12 + v15 + int32(20)
		}
	case 3:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v21 = F_strlen(m, v20)
		mBase = m.M
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		return v21 + v22 + int32(73)
	case 4:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		return v27<<(uint(int32(4))%32) - int32(-64)
	case 5:
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
		return (v34+v35)<<(uint(int32(2))%32) + int32(136)
	default:
		v48 = int32(64)
		return v48
	case 11:
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v48 = v42<<(uint(int32(2))%32) - int32(-64)
		return v48
	}
}
func F_ReorderBufferCopySnap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+168))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v16 = F_MemoryContextAllocZero(m, v8, (v9+v10)<<(uint(int32(2))%32)+int32(76))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(72)
		base.MemoryCopy(m, v16, l1, v20)
		*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = int64(1)
		v24 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+30)) = uint8(v24)
		v28 = v16 + v20
		*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		v32 = v30 << (uint(int32(2)) % 32)
		if v32 != 0 {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			base.MemoryCopy(m, v28, v33, v32)
		} else {
		}
		v35 = v28 + v32
		*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v35
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v35))) = v37
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(1)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+164))
		if v41 == int32(0) {
			v75 = int32(1)
			v77 = v35
		} else {
			v46 = l2 + int32(160)
			if v46 == v41 {
				v75 = int32(1)
				v77 = v35
			} else {
				v50 = v41
				v55 = v24
				for {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v50-int32(184))))
					*(*int32)(unsafe.Add(mBase, uint32(v56+v55<<(uint(int32(2))%32)))) = v62
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
					v65 = int32(1)
					v66 = v64 + v65
					*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v66
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
					if v70 != v46 {
						v50 = v70
						v55 = v55 + v65
						continue
					} else {
						break
					}
					break
				}
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
				v75 = v66
				v77 = v72
			}
		}
		F_pg_qsort(m, v77, v75, int32(4), int32(185))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l3
			return v16
		}
	}
}
func F_ReorderBufferForget(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l1
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v14 == v4)|base.B2i32(l1 != v14) == v4 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v44)+24)) = l2
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v47 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v21 != 0 {
		v44 = v21
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_hash_search(m, v22, v11+int32(12), int32(0), v11+int32(11))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L1
L7:
	;
	return
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	if v30 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v33
	goto L1
L10:
	;
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v37
	if v38 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v44 = v38
	goto L2
L13:
	;
	F_ReorderBufferCleanupTXN(m, l0, v44)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L28
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+172))
	if v50 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+176))
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferForget[0]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	v58 = base.B2i32(v56 != int32(0))
	goto L16
L16:
	;
	if v56 != int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_BeginInternalSubTransaction(m, int32(_a_F_ReorderBufferForget_0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v66 = int32(0)
	goto L22
L20:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_LocalExecuteInvalidationMessage(m, v53+v66<<(uint(int32(4))%32))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L24
	}
L23:
	;
	if v58 == int32(0) {
		goto L13
	} else {
		goto L26
	}
L24:
	;
	v79 = v66 + int32(1)
	if v79 != v50 {
		v66 = v79
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
	v84 = m.ExcPending
	if v84 != 0 {
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
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
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
			v47 = v13
			*(*int64)(unsafe.Add(mBase, uint32(v47)+88)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v47)+80)) = l3
			v52 = l0 + int32(12)
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v53 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v52
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v52
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v47)+100)) = v52
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
			*(*int32)(unsafe.Add(mBase, uint32(v47)+96)) = v59
			v62 = v47 + int32(96)
			*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v62
			*(*int32)(unsafe.Add(mBase, uint32(v52))) = v62
			m.G0 = v9 + int32(16)
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v20
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v23 = int32(0)
			if base.B2i32(v22 == v23)|base.B2i32(v20 != v22) == v23 {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v47 = v29
				*(*int64)(unsafe.Add(mBase, uint32(v47)+88)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v47)+80)) = l3
				v52 = l0 + int32(12)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v53 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v52
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v52
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v47)+100)) = v52
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
				*(*int32)(unsafe.Add(mBase, uint32(v47)+96)) = v59
				v62 = v47 + int32(96)
				*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v62
				*(*int32)(unsafe.Add(mBase, uint32(v52))) = v62
				m.G0 = v9 + int32(16)
				return
			} else {
				v30 = int32(0)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v37 = F_hash_search(m, v31, v9+int32(12), v30, v9+int32(11))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
					if v39 == int32(1) {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
						v43 = v42
					} else {
						v43 = v30
					}
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v44
					v47 = v43
					*(*int64)(unsafe.Add(mBase, uint32(v47)+88)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v47)+80)) = l3
					v52 = l0 + int32(12)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v53 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v52
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v47)+100)) = v52
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					*(*int32)(unsafe.Add(mBase, uint32(v47)+96)) = v59
					v62 = v47 + int32(96)
					*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v62
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v62
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v11 == v3)|base.B2i32(l1 != v11) == v3 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v18 != 0 {
			v45 = v18
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
			v50 = v47 & int32(1)
		} else {
			v50 = int32(0)
		}
		m.G0 = v8 + int32(16)
		return v50
	} else {
		v20 = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v27 = F_hash_search(m, v21, v8+int32(12), v20, v8+int32(11))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
			if v31 == int32(0) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v34
				v50 = v20
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v38
				if v39 == int32(0) {
					v50 = v20
				} else {
					v45 = v39
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
					v50 = v47 & int32(1)
				}
			}
			m.G0 = v8 + int32(16)
			return v50
		}
	}
}
func F_ReorderBufferXidSetCatalogChanges(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v6 = F_ReorderBufferTXNByXid(m, l0, l1, int32(0), l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		if v8&int32(1) == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8 | int32(1)
			v17 = l0 + int32(20)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v18 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v17
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v17
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+200)) = v17
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+196)) = v26
			v29 = v6 + int32(196)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v29
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32 + int32(1)
		} else {
		}
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
		if v37 == int32(0) {
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
			if v40&int32(1) != 0 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v37))) = v40 | int32(1)
				v47 = l0 + int32(20)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v48 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v47
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v47
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v37)+200)) = v47
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v37)+196)) = v56
				v59 = v37 + int32(196)
				*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v59
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v62 + int32(1)
			}
		}
		return
	}
}
