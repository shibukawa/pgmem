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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
		v116 = v16
		m.G0 = v14 + int32(16)
		return v116
	} else {
		if base.Ui32(v18) < base.Ui32(v17) {
			v116 = int32(1)
			m.G0 = v14 + int32(16)
			return v116
		} else {
			v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+11)))
			v23 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10)+11)))
			if v22 < v23 {
				v116 = v16
				m.G0 = v14 + int32(16)
				return v116
			} else {
				if v23 < v22 {
					v116 = int32(1)
					m.G0 = v14 + int32(16)
					return v116
				} else {
					if v22 == int32(0) {
						v30 = v9 + int32(16)
						v31 = int32(1)
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+10)))
						if v33 == v31 {
							v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+6)))
							if v38 != 0 {
								v39 = F__emscripten_memcpy_bulkmem(m, v14+int32(8), v30, v38)
								mBase = m.M
							} else {
							}
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							v42 = v41
						} else {
							v42 = v30
						}
						if v23 != 0 {
							v56 = int32(0)
						} else {
							v45 = v10 + int32(16)
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+10)))
							if v46 != int32(1) {
								v56 = v45
							} else {
								v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+6)))
								if v51 != 0 {
									v52 = F__emscripten_memcpy_bulkmem(m, v14+int32(12), v45, v51)
									mBase = m.M
								} else {
								}
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
								v56 = v54
							}
						}
						v57 = int32(36)
						v59 = v11 + v17*v57
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v59-int32(20))))
						v65 = m.T0[v64].(func(*base.Module, int32, int32, int32) int32)(m, v42, v56, v59-v57)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							if v65 < int32(0) {
								v72 = v31
							} else {
								v72 = int32(0) - v65
							}
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59-int32(28)))))
							if v75 != 0 {
								v76 = v72
							} else {
								v76 = v65
							}
							if v76 != 0 {
								v116 = v76
							} else {
								v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+6)))
								v82 = int32(17)
								v84 = int32(-2)
								v85 = (v9 + v80 + v82) & v84
								v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+6)))
								v91 = (v10 + v86 + v82) & v84
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+2)))
								v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85))))
								v97 = int32(16)
								v99 = v95 | v96<<(uint(v97)%32)
								v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+2)))
								v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91))))
								v104 = v100 | v101<<(uint(v97)%32)
								if base.Ui32(v99) < base.Ui32(v104) {
									v115 = int32(-1)
								} else {
									if base.Ui32(v104) < base.Ui32(v99) {
										v115 = int32(1)
									} else {
										v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
										v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
										if base.Ui32(v109) < base.Ui32(v110) {
											v115 = int32(-1)
										} else {
											v115 = base.B2i32(base.Ui32(v110) < base.Ui32(v109))
										}
									}
								}
								v116 = v115
							}
							m.G0 = v14 + int32(16)
							return v116
						}
					} else {
						v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+6)))
						v82 = int32(17)
						v84 = int32(-2)
						v85 = (v9 + v80 + v82) & v84
						v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+6)))
						v91 = (v10 + v86 + v82) & v84
						v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+2)))
						v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85))))
						v97 = int32(16)
						v99 = v95 | v96<<(uint(v97)%32)
						v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+2)))
						v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91))))
						v104 = v100 | v101<<(uint(v97)%32)
						if base.Ui32(v99) < base.Ui32(v104) {
							v115 = int32(-1)
						} else {
							if base.Ui32(v104) < base.Ui32(v99) {
								v115 = int32(1)
							} else {
								v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
								v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
								if base.Ui32(v109) < base.Ui32(v110) {
									v115 = int32(-1)
								} else {
									v115 = base.B2i32(base.Ui32(v110) < base.Ui32(v109))
								}
							}
						}
						v116 = v115
						m.G0 = v14 + int32(16)
						return v116
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
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
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
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
