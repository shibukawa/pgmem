package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsToastRelation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+68))
	if v5 != int32(99) {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[138]))
		v15 = base.B2i32(v10 != int32(0)) & base.B2i32(v5 == v10)
	} else {
		v15 = int32(1)
	}
	return v15
}
func F_get_toast_snapshot(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v3 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v3 != 0 {
		v18 = int32(1)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		v8 = *(*int32)(unsafe.Add(mBase, _consts[16]))
		if v8 == int32(0) {
			v18 = base.B2i32(v6 != int32(0))
		} else {
			if v6 == int32(0) {
				v18 = base.B2i32(v6 != int32(0))
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if v13 != 0 {
					v18 = base.B2i32(v6 != int32(0))
				} else {
					v18 = int32(0)
				}
			}
		}
	}
	if v18 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(92887), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(516199), int32(653), int32(92446))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		return int32(4211880)
	}
}
func F_toast_close_indexes(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v3 = int32(0)
	if v3 < l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_pfree(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L9
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+v8<<(uint(int32(2))%32))))
	F_relation_close(m, v12, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v17 = v8 + int32(1)
	if v17 != l1 {
		v8 = v17
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	return
}
func F_toast_tuple_find_biggest_attribute(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 <= int32(0) {
		return int32(-1)
	} else {
		if l1 != 0 {
			v22 = int32(48)
		} else {
			v22 = int32(16)
		}
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v35 = int32(0)
		v38 = int32(24)
		v39 = int32(-1)
		for {
			v46 = v28 + v35*int32(12)
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+8)))
			if v22&v47 != 0 {
				v77 = v38
				v78 = v39
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v35<<(uint(int32(2))%32))))
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v54 == int32(1) {
					v77 = v38
					v78 = v39
				} else {
					if v54&int32(3) == int32(2) {
						v62 = l1
					} else {
						v62 = int32(0)
					}
					if v62 != 0 {
						v77 = v38
						v78 = v39
					} else {
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v15<<(uint(int32(4))%32)+int32(20)+v35*int32(100))+84)))
						if l2 != 0 {
							if v66 != int32(109) {
								v77 = v38
								v78 = v39
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
								v72 = base.B2i32(v38 < v71)
								if v38 < v71 {
									v73 = v71
								} else {
									v73 = v38
								}
								if v38 < v71 {
									v74 = v35
								} else {
									v74 = v39
								}
								v77 = v73
								v78 = v74
							}
						} else {
							switch v66 - int32(101) {
							case 0, 19:
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
								v72 = base.B2i32(v38 < v71)
								if v38 < v71 {
									v73 = v71
								} else {
									v73 = v38
								}
								if v38 < v71 {
									v74 = v35
								} else {
									v74 = v39
								}
								v77 = v73
								v78 = v74
							default:
								v77 = v38
								v78 = v39
							}
						}
					}
				}
			}
			v80 = v35 + int32(1)
			if v80 != v15 {
				v35 = v80
				v38 = v77
				v39 = v78
				continue
			} else {
				break
			}
			break
		}
		return v78
	}
}
