package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_text_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v48 = int32(1)
			if v17&v48 != 0 {
				v52 = v48
			} else {
				v52 = int32(4)
			}
			v54 = int32(1)
			if v16&v54 != 0 {
				v58 = v54
			} else {
				v58 = int32(4)
			}
			if v16 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v88 = v75
			} else {
				v76 = int32(1)
				if v16&v76 != 0 {
					v88 = int32(base.Ui32(v16)>>(uint(v76)%32)) - v76
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v89 = F_varstr_cmp(m, v9+v52, v46, v14+v58, v88, v47)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int64(0)
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v91 != v9 {
					F_pfree(m, v9)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int64(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v95 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(int32(0) < v89))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v89))
						}
					}
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v95 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v89))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) < v89))
					}
				}
			}
		}
	}
}
func F_text_to_cstring_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
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
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	v9 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v11 == int32(1) {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if v17 == int32(18) {
				v20 = int32(16)
			} else {
				v20 = int32(0)
			}
			if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v27 = int32(4)
			} else {
				v27 = v20
			}
			v40 = int32(1)
			v41 = v27
		} else {
			if v11&int32(1) != 0 {
				v30 = int32(1)
				v40 = v11
				v41 = int32(base.Ui32(v11)>>(uint(v30)%32)) - v30
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v40 = v34
				v41 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if l2 != 0 {
			v43 = l2 - int32(1)
			if base.Ui32(v43) < base.Ui32(v41) {
				v45 = int32(1)
				if v11&v45 != 0 {
					v49 = v45
				} else {
					v49 = int32(4)
				}
				v51 = F_pg_mbcliplen(m, v9+v49, v41, v43)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
					v54 = v51
					v55 = v53
					if v54 != 0 {
						v56 = int32(1)
						if v55&v56 != 0 {
							v60 = v56
						} else {
							v60 = int32(4)
						}
						base.MemoryCopy(m, l1, v9+v60, v54)
					} else {
					}
					v64 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1+v54))) = uint8(v64)
					if l0 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				}
			} else {
				v54 = v41
				v55 = v40
				if v54 != 0 {
					v56 = int32(1)
					if v55&v56 != 0 {
						v60 = v56
					} else {
						v60 = int32(4)
					}
					base.MemoryCopy(m, l1, v9+v60, v54)
				} else {
				}
				v64 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1+v54))) = uint8(v64)
				if l0 != v9 {
					F_pfree(m, v9)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			}
		} else {
			if l0 != v9 {
				F_pfree(m, v9)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	}
}
