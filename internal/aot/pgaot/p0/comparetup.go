package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_comparetup_index_gin(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(-1)
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	if base.Ui32(v17) < base.Ui32(v18) {
		v118 = v16
		m.G0 = v14 + int32(16)
		return v118
	} else {
		if base.Ui32(v18) < base.Ui32(v17) {
			v118 = int32(1)
			m.G0 = v14 + int32(16)
			return v118
		} else {
			v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+15)))
			v23 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10)+15)))
			if v22 < v23 {
				v118 = v16
				m.G0 = v14 + int32(16)
				return v118
			} else {
				if v23 < v22 {
					v118 = int32(1)
					m.G0 = v14 + int32(16)
					return v118
				} else {
					if v22 == int32(0) {
						v30 = v9 + int32(24)
						v31 = int32(1)
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)))
						if v33 == v31 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							if v36 != 0 {
								base.MemoryCopy(m, v14+int32(8), v30, v36)
							} else {
							}
							v40 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
							v43 = v40
						} else {
							v43 = base.I64_extend_i32_u(v30)
						}
						if v23 != 0 {
							v58 = int64(0)
						} else {
							v46 = v10 + int32(24)
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
							if v47 == int32(1) {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								if v50 != 0 {
									base.MemoryCopy(m, v14+int32(8), v46, v50)
								} else {
								}
								v54 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
								v58 = v54
							} else {
								v58 = base.I64_extend_i32_u(v46)
							}
						}
						v59 = int32(36)
						v61 = v11 + v17*v59
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v61-int32(20))))
						v67 = m.T0[v66].(func(*base.Module, int64, int64, int32) int32)(m, v43, v58, v61-v59)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							if v67 < int32(0) {
								v74 = v31
							} else {
								v74 = int32(0) - v67
							}
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61-int32(28)))))
							if v77 != 0 {
								v78 = v74
							} else {
								v78 = v67
							}
							if v78 != 0 {
								v118 = v78
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								v84 = int32(25)
								v86 = int32(-2)
								v87 = (v9 + v82 + v84) & v86
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v93 = (v10 + v88 + v84) & v86
								v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+2)))
								v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
								v99 = int32(16)
								v101 = v97 | v98<<(uint(v99)%32)
								v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)))
								v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
								v106 = v102 | v103<<(uint(v99)%32)
								if base.Ui32(v101) < base.Ui32(v106) {
									v117 = int32(-1)
								} else {
									if base.Ui32(v106) < base.Ui32(v101) {
										v117 = int32(1)
									} else {
										v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+4)))
										v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
										if base.Ui32(v111) < base.Ui32(v112) {
											v117 = int32(-1)
										} else {
											v117 = base.B2i32(base.Ui32(v112) < base.Ui32(v111))
										}
									}
								}
								v118 = v117
							}
							m.G0 = v14 + int32(16)
							return v118
						}
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						v84 = int32(25)
						v86 = int32(-2)
						v87 = (v9 + v82 + v84) & v86
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						v93 = (v10 + v88 + v84) & v86
						v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+2)))
						v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
						v99 = int32(16)
						v101 = v97 | v98<<(uint(v99)%32)
						v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)))
						v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
						v106 = v102 | v103<<(uint(v99)%32)
						if base.Ui32(v101) < base.Ui32(v106) {
							v117 = int32(-1)
						} else {
							if base.Ui32(v106) < base.Ui32(v101) {
								v117 = int32(1)
							} else {
								v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+4)))
								v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
								if base.Ui32(v111) < base.Ui32(v112) {
									v117 = int32(-1)
								} else {
									v117 = base.B2i32(base.Ui32(v112) < base.Ui32(v111))
								}
							}
						}
						v118 = v117
						m.G0 = v14 + int32(16)
						return v118
					}
				}
			}
		}
	}
}
func F_comparetup_index_hash(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	v6 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v13 = v7 & v10
	if base.Ui32(v13) <= base.Ui32(v9) {
		v15 = int32(-1)
	} else {
		v15 = v11
	}
	v16 = v15 & v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v22 = v17 & v19
	if base.Ui32(v22) <= base.Ui32(v18) {
		v24 = int32(-1)
	} else {
		v24 = v20
	}
	v25 = v24 & v22
	if base.Ui32(v25) < base.Ui32(v16) {
		v61 = v6
		return v61
	} else {
		if base.Ui32(v16) < base.Ui32(v25) {
			return int32(-1)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if base.Ui32(v31) < base.Ui32(v30) {
				v61 = v6
				return v61
			} else {
				if base.Ui32(v30) < base.Ui32(v31) {
					v61 = int32(-1)
					return v61
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
					v37 = int32(16)
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+2)))
					v40 = v36<<(uint(v37)%32) | v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
					v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+2)))
					v46 = v42<<(uint(v37)%32) | v45
					if v40 != v46 {
						if base.Ui32(v40) < base.Ui32(v46) {
							v51 = int32(-1)
						} else {
							v51 = int32(1)
						}
						return v51
					} else {
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+4)))
						v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
						v61 = base.B2i32(base.Ui32(v54) < base.Ui32(v53)) - base.B2i32(base.Ui32(v53) < base.Ui32(v54))
						return v61
					}
				}
			}
		}
	}
}
