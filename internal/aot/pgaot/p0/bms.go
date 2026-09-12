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
	var v44 int32
	_ = v44
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
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
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
				F_errmsg_internal(m, int32(429946), int32(0))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(482205), int32(1040), int32(392842))
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
					v72 = int32(base.Ui32(l1) >> (uint(v66) % 32))
					if v12 == v72 {
						v102 = int32(base.Ui32(int32(-1))>>(uint(v70)%32)) & v65
					} else {
						v79 = v57 + v72<<(uint(int32(2))%32)
						v81 = v79 + int32(8)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
						*(*int32)(unsafe.Add(mBase, uint32(v81))) = v82 | v65
						if v72+int32(1) < v12 {
							v97 = F__emscripten_memset_bulkmem(m, v79+int32(12), base.I32_extend8_s(int32(255)), (v12+(v72^int32(-1)))<<(uint(int32(2))%32))
							mBase = m.M
						} else {
						}
						v102 = int32(base.Ui32(int32(-1)) >> (uint(v70) % 32))
					}
					v107 = v57 + v12<<(uint(int32(2))%32) + int32(8)
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
					*(*int32)(unsafe.Add(mBase, uint32(v107))) = v108 | v102
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
					v72 = int32(base.Ui32(l1) >> (uint(v66) % 32))
					if v12 == v72 {
						v102 = int32(base.Ui32(int32(-1))>>(uint(v70)%32)) & v65
					} else {
						v79 = v57 + v72<<(uint(int32(2))%32)
						v81 = v79 + int32(8)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
						*(*int32)(unsafe.Add(mBase, uint32(v81))) = v82 | v65
						if v72+int32(1) < v12 {
							v97 = F__emscripten_memset_bulkmem(m, v79+int32(12), base.I32_extend8_s(int32(255)), (v12+(v72^int32(-1)))<<(uint(int32(2))%32))
							mBase = m.M
						} else {
						}
						v102 = int32(base.Ui32(int32(-1)) >> (uint(v70) % 32))
					}
					v107 = v57 + v12<<(uint(int32(2))%32) + int32(8)
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
					*(*int32)(unsafe.Add(mBase, uint32(v107))) = v108 | v102
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
						v44 = v28
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v36+int32(8)+v44<<(uint(int32(2))%32)))) = int32(0)
							v54 = v44 + int32(1)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
							if v54 < v55 {
								v44 = v54
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
						v72 = int32(base.Ui32(l1) >> (uint(v66) % 32))
						if v12 == v72 {
							v102 = int32(base.Ui32(int32(-1))>>(uint(v70)%32)) & v65
						} else {
							v79 = v57 + v72<<(uint(int32(2))%32)
							v81 = v79 + int32(8)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
							*(*int32)(unsafe.Add(mBase, uint32(v81))) = v82 | v65
							if v72+int32(1) < v12 {
								v97 = F__emscripten_memset_bulkmem(m, v79+int32(12), base.I32_extend8_s(int32(255)), (v12+(v72^int32(-1)))<<(uint(int32(2))%32))
								mBase = m.M
							} else {
							}
							v102 = int32(base.Ui32(int32(-1)) >> (uint(v70) % 32))
						}
						v107 = v57 + v12<<(uint(int32(2))%32) + int32(8)
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
						*(*int32)(unsafe.Add(mBase, uint32(v107))) = v108 | v102
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
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = v7<<(uint(int32(2))%32) + int32(8)
		v12 = F_palloc(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v11 != 0 {
				v16 = F__emscripten_memcpy_bulkmem(m, v12, l0, v11)
				mBase = m.M
				v17 = v16
			} else {
				v17 = v12
			}
			return v17
		}
	}
}
func F_bms_int_members(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v126 int32
	_ = v126
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
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v22 < v23 {
				v25 = v22
			} else {
				v25 = v23
			}
			if v25 <= int32(1) {
				v28 = int32(1)
			} else {
				v28 = v25
			}
			v31 = int32(8)
			v32 = l0 + v31
			v34 = l1 + v31
			if v25 < int32(2) {
				v82 = int32(0)
				v83 = int32(-1)
			} else {
				v41 = int32(0)
				v45 = v41
				v46 = int32(-1)
				v48 = v41
				for {
					v56 = int32(2)
					v57 = v45 << (uint(v56) % 32)
					v58 = v32 + v57
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v34+v57)))
					v62 = v59 & v61
					*(*int32)(unsafe.Add(mBase, uint32(v58))) = v62
					v65 = v45 | int32(1)
					v67 = v65 << (uint(v56) % 32)
					v68 = v32 + v67
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v34+v67)))
					v72 = v69 & v71
					*(*int32)(unsafe.Add(mBase, uint32(v68))) = v72
					if v62 != 0 {
						v74 = v45
					} else {
						v74 = v46
					}
					if v72 != 0 {
						v75 = v65
					} else {
						v75 = v74
					}
					v76 = int32(2)
					v77 = v45 + v76
					v79 = v48 + v76
					if v79 != v28&int32(2147483646) {
						v45 = v77
						v46 = v75
						v48 = v79
						continue
					} else {
						break
					}
					break
				}
				v82 = v77
				v83 = v75
			}
			if v28&int32(1) != 0 {
				v94 = v82 << (uint(int32(2)) % 32)
				v95 = v32 + v94
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v34)))
				v99 = v96 & v98
				*(*int32)(unsafe.Add(mBase, uint32(v95))) = v99
				if v99 != 0 {
					v101 = v82
				} else {
					v101 = v83
				}
				v102 = v101
			} else {
				v102 = v83
			}
			if v102 == int32(-1) {
				F_pfree(m, l0)
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v102 + int32(1)
				v126 = l0
				return v126
			}
		}
	} else {
		v126 = int32(0)
		return v126
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
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v3 = int32(0)
	if v3 <= l0 {
		if l1 == int32(0) {
			v20 = v3
		} else {
			v10 = int32(base.Ui32(l0) >> (uint(int32(5)) % 32))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v11 <= v10 {
				v20 = v3
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1+v10<<(uint(int32(2))%32))+8))
				v20 = int32(base.Ui32(v16)>>(uint(l0)%32)) & int32(1)
			}
		}
		return v20
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(429946), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(482205), int32(519), int32(224017))
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
	var v48 int32
	_ = v48
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
		v48 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v48
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 < v13 {
		v48 = v3
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+(l1+v20))))
	v39 = v34 & (v36 ^ int32(-1))
	v41 = base.B2i32(v39 == int32(0))
	if v39 != 0 {
		v48 = v41
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v48 = v41
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
