package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_text_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v10 + int32(1)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v25 = v23 & int32(1)
			if v25 != 0 {
				v26 = v15
			} else {
				v26 = v10 + int32(4)
			}
			if v23 == int32(1) {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v32 == int32(18) {
					v35 = int32(16)
				} else {
					v35 = int32(0)
				}
				if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v42 = int32(4)
				} else {
					v42 = v35
				}
				v53 = v42
			} else {
				v43 = int32(1)
				if v25 != 0 {
					v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v54 = int32(1)
			v55 = v17 + v54
			if v19&v54 != 0 {
				v60 = v55
			} else {
				v60 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
				if v66 == int32(18) {
					v69 = int32(16)
				} else {
					v69 = int32(0)
				}
				if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v76 = int32(4)
				} else {
					v76 = v69
				}
				v89 = v76
			} else {
				v77 = int32(1)
				if v19&v77 != 0 {
					v89 = int32(base.Ui32(v19)>>(uint(v77)%32)) - v77
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_varstr_cmp(m, v26, v53, v60, v89, v20)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								return base.B2i32(int32(0) < v90)
							}
						} else {
							return base.B2i32(int32(0) < v90)
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v90)
						}
					} else {
						return base.B2i32(int32(0) < v90)
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
