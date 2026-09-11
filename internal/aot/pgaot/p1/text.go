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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
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
				v29 = int32(4)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v31&int32(254) == int32(2) {
					v40 = v29
				} else {
					v40 = base.B2i32(v31 == int32(18)) << (uint(v29) % 32)
				}
				if v31 == int32(1) {
					v43 = v29
				} else {
					v43 = v40
				}
				v54 = v43
			} else {
				v44 = int32(1)
				if v25 != 0 {
					v54 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v55 = int32(1)
			v56 = v17 + v55
			if v19&v55 != 0 {
				v61 = v56
			} else {
				v61 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v64 = int32(4)
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
				if v66&int32(254) == int32(2) {
					v75 = v64
				} else {
					v75 = base.B2i32(v66 == int32(18)) << (uint(v64) % 32)
				}
				if v66 == int32(1) {
					v78 = v64
				} else {
					v78 = v75
				}
				v91 = v78
			} else {
				v79 = int32(1)
				if v19&v79 != 0 {
					v91 = int32(base.Ui32(v19)>>(uint(v79)%32)) - v79
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v91 = int32(base.Ui32(v85)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v92 = F_varstr_cmp(m, v26, v54, v61, v91, v20)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v94 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v98 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								return base.B2i32(int32(0) < v92)
							}
						} else {
							return base.B2i32(int32(0) < v92)
						}
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v98 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v92)
						}
					} else {
						return base.B2i32(int32(0) < v92)
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	v9 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v11 == int32(1) {
			v14 = int32(4)
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if v16&int32(254) == int32(2) {
				v25 = v14
			} else {
				v25 = base.B2i32(v16 == int32(18)) << (uint(v14) % 32)
			}
			if v16 == int32(1) {
				v28 = v14
			} else {
				v28 = v25
			}
			v41 = int32(1)
			v42 = v28
		} else {
			if v11&int32(1) != 0 {
				v31 = int32(1)
				v41 = v11
				v42 = int32(base.Ui32(v11)>>(uint(v31)%32)) - v31
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v41 = v35
				v42 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if l2 != 0 {
			v43 = int32(1)
			v46 = l2 - v43
			if base.Ui32(v46) < base.Ui32(v42) {
				v48 = int32(1)
				if v11&v48 != 0 {
					v52 = v48
				} else {
					v52 = int32(4)
				}
				v54 = F_pg_mbcliplen(m, v9+v52, v42, v46)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
					v57 = v54
					v58 = v56
					if v58&int32(1) != 0 {
						v61 = v43
					} else {
						v61 = int32(4)
					}
					if v57 != 0 {
						v63 = F__emscripten_memcpy_bulkmem(m, l1, v9+v61, v57)
						mBase = m.M
						v64 = v63
					} else {
						v64 = l1
					}
					v66 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v64+v57))) = uint8(v66)
					if l0 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				}
			} else {
				v57 = v42
				v58 = v41
				if v58&int32(1) != 0 {
					v61 = v43
				} else {
					v61 = int32(4)
				}
				if v57 != 0 {
					v63 = F__emscripten_memcpy_bulkmem(m, l1, v9+v61, v57)
					mBase = m.M
					v64 = v63
				} else {
					v64 = l1
				}
				v66 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v64+v57))) = uint8(v66)
				if l0 != v9 {
					F_pfree(m, v9)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
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
				v72 = m.ExcPending
				if v72 != 0 {
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
