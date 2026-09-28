package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datumRestore(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6 + int32(4)
	if v7 == int32(-2) {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v13)
		return int64(0)
	} else {
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v17)
		if v7 == int32(-1) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v21 + int32(8)
			return v22
		} else {
			v27 = F_palloc(m, v7)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				if v7 != 0 {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					base.MemoryCopy(m, v27, v31, v7)
				} else {
				}
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v33 + v7
				return base.I64_extend_i32_u(v27)
			}
		}
	}
}
func F_datumSerialize(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	if l1 == int32(0) {
		if l2 == int32(0) {
			v11 = base.I32_wrap_i64(l0)
			if l3 != int32(-1) {
				v23 = F_datumGetSize(m, l0, int32(0), l3)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = v23
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					v29 = v27 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v29
					v73 = v23
					v74 = v29
					if v73 != 0 {
						base.MemoryCopy(m, v74, v11, v73)
					} else {
					}
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v77 + v73
					return
				}
			} else {
				v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				if v14 != int32(1) {
					v23 = F_datumGetSize(m, l0, int32(0), l3)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = v23
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						v29 = v27 + int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v29
						v73 = v23
						v74 = v29
						if v73 != 0 {
							base.MemoryCopy(m, v74, v11, v73)
						} else {
						}
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v77 + v73
						return
					}
				} else {
					v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
					if v17&int32(254) == int32(2) {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+2))
						v45 = F_EOH_get_flat_size(m, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(v47))) = v45
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							v51 = v49 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
							if v44 == int32(0) {
								v73 = v45
								v74 = v51
								if v73 != 0 {
									base.MemoryCopy(m, v74, v11, v73)
								} else {
								}
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v77 + v73
								return
							} else {
								v55 = F_palloc(m, v45)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_EOH_flatten_into(m, v44, v55, v45)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										if v45 != 0 {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
											base.MemoryCopy(m, v59, v55, v45)
										} else {
										}
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v61 + v45
										F_pfree(m, v55)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v23 = F_datumGetSize(m, l0, int32(0), l3)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v23
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							v29 = v27 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v29
							v73 = v23
							v74 = v29
							if v73 != 0 {
								base.MemoryCopy(m, v74, v11, v73)
							} else {
							}
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v77 + v73
							return
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(-1)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v34 + int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v34)+4)) = l0
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v39 + int32(8)
			return
		}
	} else {
		v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(-2)
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v69 + int32(4)
		return
	}
}
func F_datum_compute_size(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = base.B2i32(l3 != int32(-1))
	if v13|base.B2i32(l4 == int32(112)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v117
L2:
	;
	v117 = v106 + v108
	goto L1
L3:
	;
	v103 = F_strlen(m, v73)
	mBase = m.M
	v106 = v70
	v108 = v103 + int32(1)
	goto L2
L4:
	;
	if v79 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	switch l2 - int32(99) {
	case 0:
		v68 = l0
		v69 = int32(-1)
		goto L14
	case 1:
		goto L15
	default:
		goto L18
	case 6:
		goto L16
	case 16:
		goto L17
	}
L6:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v36&int32(1) != 0 {
		v77 = l0
		v78 = v34
		v79 = v36
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v19 = base.I32_wrap_i64(l1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v20&int32(3) != 0 {
		v34 = v19
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if l3 != int32(-1) {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v25 = int32(base.Ui32(v23) >> (uint(int32(2)) % 32))
	if base.Ui32(int32(127)) < base.Ui32(v25-int32(3)) {
		v34 = v19
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v117 = l0 + v25 - int32(3)
	goto L1
L12:
	;
	v34 = base.I32_wrap_i64(l1)
	goto L6
L13:
	;
	goto L5
L14:
	;
	v70 = v68 & v69
	if int32(0) < l3 {
		v106 = v70
		v108 = l3
		goto L2
	} else {
		goto L23
	}
L15:
	;
	v68 = l0 + int32(7)
	v69 = int32(-8)
	goto L14
L16:
	;
	v68 = l0 + int32(3)
	v69 = int32(-4)
	goto L14
L17:
	;
	v68 = l0 + int32(1)
	v69 = int32(-2)
	goto L14
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
	F_errmsg_internal(m, int32(_a_F_datum_compute_size_0), v10)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_datum_compute_size_1), int32(322), int32(_a_F_datum_compute_size_2))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v73 = base.I32_wrap_i64(l1)
	if l3 != int32(-1) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v77 = v70
	v78 = v73
	v79 = v76
	goto L4
L25:
	;
	v83 = int32(18)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v85 == v83 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v79&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v88 = v83
	goto L30
L29:
	;
	v88 = int32(2)
	goto L30
L30:
	;
	if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v95 = int32(6)
	goto L33
L32:
	;
	v95 = v88
	goto L33
L33:
	;
	v106 = v77
	v108 = v95
	goto L2
L34:
	;
	v106 = v77
	v108 = int32(base.Ui32(v79) >> (uint(int32(1)) % 32))
	goto L2
L35:
	;
	goto L36
L36:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v106 = v77
	v108 = int32(base.Ui32(v100) >> (uint(int32(2)) % 32))
	goto L2
}
