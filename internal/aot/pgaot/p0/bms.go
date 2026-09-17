package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bms_add_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	if l1 <= l2 {
		if l1 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_bms_add_range_0), int32(0))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_bms_add_range_1), int32(1040), int32(_a_F_bms_add_range_2))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v12 = base.I32_div_s(l2, int32(32))
			if l0 == int32(0) {
				v16 = v12 + int32(1)
				v21 = F_palloc0(m, v16<<(uint(int32(2))%32)+int32(8))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v16
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(445)
					v57 = v21
					v65 = int32(-1) << (uint(l1) % 32)
					v66 = int32(5)
					v70 = v12<<(uint(v66)%32) - l2 + int32(31)
					v75 = v57 + v12<<(uint(int32(2))%32) + int32(8)
					v77 = int32(base.Ui32(l1) >> (uint(v66) % 32))
					if v12 == v77 {
						v107 = int32(base.Ui32(int32(-1))>>(uint(v70)%32)) & v65
					} else {
						v84 = v57 + v77<<(uint(int32(2))%32)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v85 | v65
						if v12 <= v77+int32(1) {
						} else {
							v95 = (v12 + (v77 ^ int32(-1))) << (uint(int32(2)) % 32)
							if v95 == int32(0) {
							} else {
								base.MemoryFill(m, v84+int32(12), int32(255), v95)
							}
						}
						v107 = int32(base.Ui32(int32(-1)) >> (uint(v70) % 32))
					}
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
					*(*int32)(unsafe.Add(mBase, uint32(v75))) = v107 | v108
					v111 = v57
					return v111
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v12 < v28 {
					v57 = l0
					v65 = int32(-1) << (uint(l1) % 32)
					v66 = int32(5)
					v70 = v12<<(uint(v66)%32) - l2 + int32(31)
					v75 = v57 + v12<<(uint(int32(2))%32) + int32(8)
					v77 = int32(base.Ui32(l1) >> (uint(v66) % 32))
					if v12 == v77 {
						v107 = int32(base.Ui32(int32(-1))>>(uint(v70)%32)) & v65
					} else {
						v84 = v57 + v77<<(uint(int32(2))%32)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v85 | v65
						if v12 <= v77+int32(1) {
						} else {
							v95 = (v12 + (v77 ^ int32(-1))) << (uint(int32(2)) % 32)
							if v95 == int32(0) {
							} else {
								base.MemoryFill(m, v84+int32(12), int32(255), v95)
							}
						}
						v107 = int32(base.Ui32(int32(-1)) >> (uint(v70) % 32))
					}
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
					*(*int32)(unsafe.Add(mBase, uint32(v75))) = v107 | v108
					v111 = v57
					return v111
				} else {
					v31 = v12 + int32(1)
					v36 = F_repalloc(m, l0, v31<<(uint(int32(2))%32)+int32(8))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v31
						v45 = v28
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v36+int32(8)+v45<<(uint(int32(2))%32)))) = int32(0)
							v54 = v45 + int32(1)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
							if v54 < v55 {
								v45 = v54
								continue
							} else {
								break
							}
							break
						}
						v57 = v36
						v65 = int32(-1) << (uint(l1) % 32)
						v66 = int32(5)
						v70 = v12<<(uint(v66)%32) - l2 + int32(31)
						v75 = v57 + v12<<(uint(int32(2))%32) + int32(8)
						v77 = int32(base.Ui32(l1) >> (uint(v66) % 32))
						if v12 == v77 {
							v107 = int32(base.Ui32(int32(-1))>>(uint(v70)%32)) & v65
						} else {
							v84 = v57 + v77<<(uint(int32(2))%32)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v85 | v65
							if v12 <= v77+int32(1) {
							} else {
								v95 = (v12 + (v77 ^ int32(-1))) << (uint(int32(2)) % 32)
								if v95 == int32(0) {
								} else {
									base.MemoryFill(m, v84+int32(12), int32(255), v95)
								}
							}
							v107 = int32(base.Ui32(int32(-1)) >> (uint(v70) % 32))
						}
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
						*(*int32)(unsafe.Add(mBase, uint32(v75))) = v107 | v108
						v111 = v57
						return v111
					}
				}
			}
		}
	} else {
		v111 = l0
		return v111
	}
}
func F_bms_copy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = v8<<(uint(int32(2))%32) + int32(8)
		v13 = F_palloc(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				base.MemoryCopy(m, v13, l0, v12)
			} else {
			}
			return v13
		}
	}
}
func F_bms_int_members(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v137 int32
	_ = v137
	if l0 != 0 {
		if l1 == int32(0) {
			F_pfree(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			v21 = int32(8)
			v22 = l0 + v21
			v24 = l1 + v21
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v25 < v26 {
				v28 = v25
			} else {
				v28 = v26
			}
			if v28 < int32(2) {
				v83 = int32(0)
				v84 = int32(-1)
				v95 = v83 << (uint(int32(2)) % 32)
				v96 = v22 + v95
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v24+v95)))
				v100 = v97 & v99
				*(*int32)(unsafe.Add(mBase, uint32(v96))) = v100
				if v100 != 0 {
					v102 = v83
				} else {
					v102 = v84
				}
				v105 = v102
			} else {
				v33 = int32(1)
				if v28 <= v33 {
					v36 = v33
				} else {
					v36 = v28
				}
				v44 = int32(0)
				v45 = int32(-1)
				v54 = int32(0)
				for {
					v55 = int32(2)
					v56 = v44 << (uint(v55) % 32)
					v57 = v22 + v56
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v24+v56)))
					v61 = v58 & v60
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = v61
					v64 = v44 | int32(1)
					v66 = v64 << (uint(v55) % 32)
					v67 = v22 + v66
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v24+v66)))
					v71 = v68 & v70
					*(*int32)(unsafe.Add(mBase, uint32(v67))) = v71
					if v61 != 0 {
						v73 = v44
					} else {
						v73 = v45
					}
					if v71 != 0 {
						v74 = v64
					} else {
						v74 = v73
					}
					v75 = int32(2)
					v76 = v44 + v75
					v78 = v54 + v75
					if v78 != v36&int32(2147483646) {
						v44 = v76
						v45 = v74
						v54 = v78
						continue
					} else {
						break
					}
					break
				}
				if v36&int32(1) == int32(0) {
					v105 = v74
				} else {
					v83 = v76
					v84 = v74
					v95 = v83 << (uint(int32(2)) % 32)
					v96 = v22 + v95
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v24+v95)))
					v100 = v97 & v99
					*(*int32)(unsafe.Add(mBase, uint32(v96))) = v100
					if v100 != 0 {
						v102 = v83
					} else {
						v102 = v84
					}
					v105 = v102
				}
			}
			if v105 == int32(-1) {
				F_pfree(m, l0)
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105 + int32(1)
				v137 = l0
				return v137
			}
		}
	} else {
		v137 = int32(0)
		return v137
	}
}
func F_bms_is_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v3 = int32(0)
	if v3 <= l0 {
		if l1 == int32(0) {
			v21 = v3
		} else {
			v10 = int32(base.Ui32(l0) >> (uint(int32(5)) % 32))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v11 <= v10 {
				v21 = v3
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1+v10<<(uint(int32(2))%32))+8))
				v21 = int32(base.Ui32(v16)>>(uint(l0)%32)) & int32(1)
			}
		}
		return v21
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_bms_is_member_0), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_bms_is_member_1), int32(519), int32(_a_F_bms_is_member_2))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_bms_is_subset(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	if l1 == int32(0) {
		v50 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v50
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 < v13 {
		v50 = v3
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v16 = int32(1)
	if v13 <= v16 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v19 = v16
	goto L9
L8:
	;
	v19 = v13
	goto L9
L9:
	;
	v20 = int32(8)
	v25 = int32(0)
	goto L10
L10:
	;
	v32 = v25 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0+v20+v32)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+v20+v32)))
	v39 = v34 & (v36 ^ int32(-1))
	v41 = base.B2i32(v39 == int32(0))
	if v39 != 0 {
		v50 = v41
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v50 = v41
	goto L4
L12:
	;
	v43 = v25 + int32(1)
	if v43 != v19 {
		v25 = v43
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
