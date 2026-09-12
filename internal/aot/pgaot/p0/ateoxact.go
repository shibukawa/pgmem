package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_ApplyLauncher(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	if l0 == int32(0) {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[295])) = uint8(v18)
		return
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[295])))
		if v5 == int32(0) {
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _consts[295])) = uint8(v18)
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, _consts[514]))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			if v10 == int32(0) {
				v18 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _consts[295])) = uint8(v18)
				return
			} else {
				v14 = F_kill(m, v10, int32(10))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v18 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _consts[295])) = uint8(v18)
					return
				}
			}
		}
	}
}
func F_AtEOXact_RelationMap(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v4 = m.G0
	v6 = v4 - int32(528)
	m.G0 = v6
	if l1 != 0 {
		v123 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[115])) = v123
		*(*int32)(unsafe.Add(mBase, _consts[114])) = v123
		*(*int32)(unsafe.Add(mBase, _consts[116])) = v123
		*(*int32)(unsafe.Add(mBase, _consts[117])) = v123
		m.G0 = v6 + int32(528)
		return
	} else {
		if l0 == int32(0) {
			v123 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[115])) = v123
			*(*int32)(unsafe.Add(mBase, _consts[114])) = v123
			*(*int32)(unsafe.Add(mBase, _consts[116])) = v123
			*(*int32)(unsafe.Add(mBase, _consts[117])) = v123
			m.G0 = v6 + int32(528)
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[114]))
			if v11 != 0 {
				v13 = *(*int32)(unsafe.Add(mBase, _consts[29]))
				v17 = F_LWLockAcquire(m, v13+int32(3200), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_read_relmap_file(m, int32(4475724), int32(312550), int32(1), int32(22))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v29 = F__emscripten_memcpy_bulkmem(m, v6+int32(4), int32(4475724), int32(524))
						mBase = m.M
						v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
						F_merge_map_updates(m, v6+int32(4), int32(4475200), v35)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v40 = int32(1)
							F_write_relmap_file(m, v6+int32(4), v40, v40, v40, int32(0), int32(1664), int32(312550))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								v52 = F__emscripten_memcpy_bulkmem(m, int32(4475724), v6+int32(4), int32(524))
								mBase = m.M
								v55 = *(*int32)(unsafe.Add(mBase, _consts[29]))
								F_LWLockRelease(m, v55+int32(3200))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[114])) = int32(0)
									v64 = *(*int32)(unsafe.Add(mBase, _consts[115]))
									if v64 == int32(0) {
										m.G0 = v6 + int32(528)
										return
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, _consts[29]))
										v72 = F_LWLockAcquire(m, v68+int32(3200), int32(0))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, _consts[883]))
											F_read_relmap_file(m, int32(4476772), v76, int32(1), int32(22))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												v85 = F__emscripten_memcpy_bulkmem(m, v6+int32(4), int32(4476772), int32(524))
												mBase = m.M
												v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
												F_merge_map_updates(m, v6+int32(4), int32(4476248), v91)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													v96 = int32(1)
													v100 = *(*int32)(unsafe.Add(mBase, _consts[226]))
													v102 = *(*int32)(unsafe.Add(mBase, _consts[271]))
													v104 = *(*int32)(unsafe.Add(mBase, _consts[883]))
													F_write_relmap_file(m, v6+int32(4), v96, v96, v96, v100, v102, v104)
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return
													} else {
														v111 = F__emscripten_memcpy_bulkmem(m, int32(4476772), v6+int32(4), int32(524))
														mBase = m.M
														v114 = *(*int32)(unsafe.Add(mBase, _consts[29]))
														F_LWLockRelease(m, v114+int32(3200))
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[115])) = int32(0)
															m.G0 = v6 + int32(528)
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
					}
				}
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, _consts[115]))
				if v64 == int32(0) {
					m.G0 = v6 + int32(528)
					return
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, _consts[29]))
					v72 = F_LWLockAcquire(m, v68+int32(3200), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, _consts[883]))
						F_read_relmap_file(m, int32(4476772), v76, int32(1), int32(22))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v85 = F__emscripten_memcpy_bulkmem(m, v6+int32(4), int32(4476772), int32(524))
							mBase = m.M
							v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
							F_merge_map_updates(m, v6+int32(4), int32(4476248), v91)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								v96 = int32(1)
								v100 = *(*int32)(unsafe.Add(mBase, _consts[226]))
								v102 = *(*int32)(unsafe.Add(mBase, _consts[271]))
								v104 = *(*int32)(unsafe.Add(mBase, _consts[883]))
								F_write_relmap_file(m, v6+int32(4), v96, v96, v96, v100, v102, v104)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									v111 = F__emscripten_memcpy_bulkmem(m, int32(4476772), v6+int32(4), int32(524))
									mBase = m.M
									v114 = *(*int32)(unsafe.Add(mBase, _consts[29]))
									F_LWLockRelease(m, v114+int32(3200))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[115])) = int32(0)
										m.G0 = v6 + int32(528)
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
}
func F_AtEOXact_SPI(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v7 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	if v7 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[348]))
	v14 = v11 + v7<<(uint(int32(6))%32)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+41)))
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	*(*int64)(unsafe.Add(mBase, _consts[349])) = v17
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	*(*int32)(unsafe.Add(mBase, _consts[350])) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v25 = v7 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[347])) = v25
	*(*int32)(unsafe.Add(mBase, _consts[351])) = v22
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = v11 + v25<<(uint(int32(6))%32)
	goto L6
L5:
	;
	v34 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[346])) = v34
	if v7 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if l0 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L8:
	;
	v40 = v25
	goto L9
L9:
	;
	v45 = v11 + v40<<(uint(int32(6))%32)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+41)))
	if v46 != 0 {
		goto L7
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	*(*int64)(unsafe.Add(mBase, _consts[349])) = v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	*(*int32)(unsafe.Add(mBase, _consts[350])) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+60))
	v56 = v40 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[347])) = v56
	*(*int32)(unsafe.Add(mBase, _consts[351])) = v53
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v65 = v11 + v56<<(uint(int32(6))%32)
	goto L14
L13:
	;
	v65 = int32(0)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[346])) = v65
	if int32(0) < v40 {
		v40 = v56
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	v78 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	if v78 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(64))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(316608), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	F_errhint(m, int32(575956), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(494534), int32(472), int32(530812))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L1
}
