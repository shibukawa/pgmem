package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplesort_begin_batch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v6 = int32(4515600)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v12&int32(2) == int32(0) {
		v19 = F_BumpContextCreate(m, v11, int32(164483), int32(8388608))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v27 = v19
			v28 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v28)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v27
			v33 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+144)) = uint16(v33)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v28
			v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v39
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			if v42 != 0 {
				if v41 == int32(1024) {
					v65 = v39
					if int64(0) <= v65 {
						v82 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
						return
					} else {
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
						if v68 != 0 {
							v82 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
							*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(78556), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									F_errfinish(m, int32(493093), int32(811), int32(325991))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					F_pfree(m, v42)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(1024)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
						v54 = int32(16384)
						v55 = F_palloc(m, v54)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v55
							v58 = F_GetMemoryChunkSpace(m, v55)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
								v62 = v60 - base.I64_extend_i32_u(v58)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v62
								v65 = v62
								if int64(0) <= v65 {
									v82 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
									*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
									*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
									return
								} else {
									v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
									if v68 != 0 {
										v82 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
										*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(78556), int32(0))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												F_errfinish(m, int32(493093), int32(811), int32(325991))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
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
				v54 = v41 << (uint(int32(4)) % 32)
				v55 = F_palloc(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v55
					v58 = F_GetMemoryChunkSpace(m, v55)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
						v62 = v60 - base.I64_extend_i32_u(v58)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v62
						v65 = v62
						if int64(0) <= v65 {
							v82 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
							*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
							return
						} else {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
							if v68 != 0 {
								v82 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
								*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(78556), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										F_errfinish(m, int32(493093), int32(811), int32(325991))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
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
		v25 = F_AllocSetContextCreateInternal(m, v11, int32(164483), int32(0), int32(8192), int32(8388608))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = v25
			v28 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v28)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v27
			v33 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+144)) = uint16(v33)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v28
			v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v39
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			if v42 != 0 {
				if v41 == int32(1024) {
					v65 = v39
					if int64(0) <= v65 {
						v82 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
						return
					} else {
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
						if v68 != 0 {
							v82 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
							*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(78556), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									F_errfinish(m, int32(493093), int32(811), int32(325991))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					F_pfree(m, v42)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(1024)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
						v54 = int32(16384)
						v55 = F_palloc(m, v54)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v55
							v58 = F_GetMemoryChunkSpace(m, v55)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
								v62 = v60 - base.I64_extend_i32_u(v58)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v62
								v65 = v62
								if int64(0) <= v65 {
									v82 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
									*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
									*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
									return
								} else {
									v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
									if v68 != 0 {
										v82 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
										*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(78556), int32(0))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												F_errfinish(m, int32(493093), int32(811), int32(325991))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
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
				v54 = v41 << (uint(int32(4)) % 32)
				v55 = F_palloc(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v55
					v58 = F_GetMemoryChunkSpace(m, v55)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
						v62 = v60 - base.I64_extend_i32_u(v58)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v62
						v65 = v62
						if int64(0) <= v65 {
							v82 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
							*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
							return
						} else {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
							if v68 != 0 {
								v82 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v82
								*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v82
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(78556), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										F_errfinish(m, int32(493093), int32(811), int32(325991))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
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
}
func F_tuplesort_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = F_tuplesort_begin_common(m, l6, int32(0), l7)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = int32(4515600)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v25
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
	if v28 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(1836)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = l0
	v57 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+36)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(1837)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(1838)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(1839)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(1840)
	v69 = F_palloc0(m, l1*int32(36))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L12
	}
L4:
	;
	v33 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v33 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	if l7&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v43 = int32(116)
	goto L9
L8:
	;
	v43 = int32(102)
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v43
	F_errmsg_internal(m, int32(502246), v15)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(493887), int32(199), int32(238917))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L3
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v69
	if l1 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v23
	m.G0 = v15 + int32(16)
	return v18
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+9)) = uint8(v79)
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	*(*uint16)(unsafe.Add(mBase, uint32(v69)+10)) = uint16(v81)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+20)) = uint8(v83)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_PrepareSortSupportFromOrderingOp(m, v85, v69)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v88 = int32(1)
	if l1 != v88 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v98 = v88
	goto L19
L17:
	;
	goto L18
L18:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+24))
	if v133 != 0 {
		goto L13
	} else {
		goto L23
	}
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v106 = v103 + v98*int32(36)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v108
	v111 = v98 << (uint(int32(2)) % 32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l4+v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v113
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v98))))
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+9)) = uint8(v116)
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v98<<(uint(int32(1))%32)))))
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+20)) = uint8(v122)
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+10)) = uint16(v121)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v111+l3)))
	F_PrepareSortSupportFromOrderingOp(m, v126, v106)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v130 = v98 + int32(1)
	if v130 != l1 {
		v98 = v130
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L13
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v132
	goto L13
}
func F_tuplesort_begin_index_brin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_tuplesort_begin_common(m, l0, l1, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
		if v14 != int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1854)
			v36 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v36
			*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)) = uint8(v36)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(1855)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(1856)
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1857)
			m.G0 = v6 + int32(16)
			return v9
		} else {
			v19 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1854)
					v36 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)) = uint8(v36)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(1855)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(1856)
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1857)
					m.G0 = v6 + int32(16)
					return v9
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(102)
					F_errmsg_internal(m, int32(502041), v6)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493887), int32(567), int32(275647))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1854)
							v36 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v36
							*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)) = uint8(v36)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(1855)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(1856)
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1857)
							m.G0 = v6 + int32(16)
							return v9
						}
					}
				}
			}
		}
	}
}
func F_tuplesort_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = int32(4515600)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v16 == int32(0) {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v25 = base.I64_div_s(v19-v20+int64(1023), int64(1024))
		v35 = v25
		v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
		if v37 != int32(1) {
			v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v94 != 0 {
				m.T0[v94].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					F_MemoryContextReset(m, v99)
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return
					} else {
						m.G0 = v9 - int32(-64)
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				F_MemoryContextReset(m, v99)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					m.G0 = v9 - int32(-64)
					return
				}
			}
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v43 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				if v40 != 0 {
					if v43 == int32(0) {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v94 != 0 {
							m.T0[v94].(func(*base.Module, int32))(m, l0)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_MemoryContextReset(m, v99)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									m.G0 = v9 - int32(-64)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							F_MemoryContextReset(m, v99)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								m.G0 = v9 - int32(-64)
								return
							}
						}
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
						v52 = F_pg_rusage_show(m, l0+int32(256))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v52
							*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v49
							if v48 != 0 {
								v59 = int32(78751)
							} else {
								v59 = int32(78760)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v59
							F_errmsg_internal(m, int32(203933), v7+int32(-32))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v88 = int32(923)
								F_errfinish(m, int32(493093), v88, int32(410716))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v94 != 0 {
										m.T0[v94].(func(*base.Module, int32))(m, l0)
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
											v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											F_MemoryContextReset(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												m.G0 = v9 - int32(-64)
												return
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
										v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										F_MemoryContextReset(m, v99)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											m.G0 = v9 - int32(-64)
											return
										}
									}
								}
							}
						}
					}
				} else {
					if v43 == int32(0) {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v94 != 0 {
							m.T0[v94].(func(*base.Module, int32))(m, l0)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_MemoryContextReset(m, v99)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									m.G0 = v9 - int32(-64)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							F_MemoryContextReset(m, v99)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								m.G0 = v9 - int32(-64)
								return
							}
						}
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
						v73 = F_pg_rusage_show(m, l0+int32(256))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v70
							if v69 != 0 {
								v80 = int32(78674)
							} else {
								v80 = int32(78774)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v80
							F_errmsg_internal(m, int32(203982), v9)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								v88 = int32(927)
								F_errfinish(m, int32(493093), v88, int32(410716))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v94 != 0 {
										m.T0[v94].(func(*base.Module, int32))(m, l0)
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
											v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											F_MemoryContextReset(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												m.G0 = v9 - int32(-64)
												return
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
										v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										F_MemoryContextReset(m, v99)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											m.G0 = v9 - int32(-64)
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
	} else {
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
		v28 = v26 - v27
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		if v29 == int32(0) {
			v35 = v28
			v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
			if v37 != int32(1) {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v94 != 0 {
					m.T0[v94].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						F_MemoryContextReset(m, v99)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							m.G0 = v9 - int32(-64)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					F_MemoryContextReset(m, v99)
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return
					} else {
						m.G0 = v9 - int32(-64)
						return
					}
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v43 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					if v40 != 0 {
						if v43 == int32(0) {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v94 != 0 {
								m.T0[v94].(func(*base.Module, int32))(m, l0)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									F_MemoryContextReset(m, v99)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										m.G0 = v9 - int32(-64)
										return
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_MemoryContextReset(m, v99)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									m.G0 = v9 - int32(-64)
									return
								}
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
							v52 = F_pg_rusage_show(m, l0+int32(256))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v52
								*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v35
								*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v49
								if v48 != 0 {
									v59 = int32(78751)
								} else {
									v59 = int32(78760)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v59
								F_errmsg_internal(m, int32(203933), v7+int32(-32))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									v88 = int32(923)
									F_errfinish(m, int32(493093), v88, int32(410716))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v94 != 0 {
											m.T0[v94].(func(*base.Module, int32))(m, l0)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
												v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												F_MemoryContextReset(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													m.G0 = v9 - int32(-64)
													return
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
											v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											F_MemoryContextReset(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												m.G0 = v9 - int32(-64)
												return
											}
										}
									}
								}
							}
						}
					} else {
						if v43 == int32(0) {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v94 != 0 {
								m.T0[v94].(func(*base.Module, int32))(m, l0)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									F_MemoryContextReset(m, v99)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										m.G0 = v9 - int32(-64)
										return
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_MemoryContextReset(m, v99)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									m.G0 = v9 - int32(-64)
									return
								}
							}
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
							v73 = F_pg_rusage_show(m, l0+int32(256))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v73
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v35
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v70
								if v69 != 0 {
									v80 = int32(78674)
								} else {
									v80 = int32(78774)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v80
								F_errmsg_internal(m, int32(203982), v9)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									v88 = int32(927)
									F_errfinish(m, int32(493093), v88, int32(410716))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v94 != 0 {
											m.T0[v94].(func(*base.Module, int32))(m, l0)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
												v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												F_MemoryContextReset(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													m.G0 = v9 - int32(-64)
													return
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
											v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											F_MemoryContextReset(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												m.G0 = v9 - int32(-64)
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
		} else {
			F_LogicalTapeSetClose(m, v29)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v35 = v28
				v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
				if v37 != int32(1) {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v94 != 0 {
						m.T0[v94].(func(*base.Module, int32))(m, l0)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							F_MemoryContextReset(m, v99)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								m.G0 = v9 - int32(-64)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						F_MemoryContextReset(m, v99)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							m.G0 = v9 - int32(-64)
							return
						}
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v43 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						if v40 != 0 {
							if v43 == int32(0) {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v94 != 0 {
									m.T0[v94].(func(*base.Module, int32))(m, l0)
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
										v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										F_MemoryContextReset(m, v99)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											m.G0 = v9 - int32(-64)
											return
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									F_MemoryContextReset(m, v99)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										m.G0 = v9 - int32(-64)
										return
									}
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
								v52 = F_pg_rusage_show(m, l0+int32(256))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v52
									*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v35
									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v49
									if v48 != 0 {
										v59 = int32(78751)
									} else {
										v59 = int32(78760)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v59
									F_errmsg_internal(m, int32(203933), v7+int32(-32))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										v88 = int32(923)
										F_errfinish(m, int32(493093), v88, int32(410716))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v94 != 0 {
												m.T0[v94].(func(*base.Module, int32))(m, l0)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
													v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													F_MemoryContextReset(m, v99)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return
													} else {
														m.G0 = v9 - int32(-64)
														return
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
												v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												F_MemoryContextReset(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													m.G0 = v9 - int32(-64)
													return
												}
											}
										}
									}
								}
							}
						} else {
							if v43 == int32(0) {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v94 != 0 {
									m.T0[v94].(func(*base.Module, int32))(m, l0)
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
										v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										F_MemoryContextReset(m, v99)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											m.G0 = v9 - int32(-64)
											return
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									F_MemoryContextReset(m, v99)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										m.G0 = v9 - int32(-64)
										return
									}
								}
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
								v73 = F_pg_rusage_show(m, l0+int32(256))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v73
									*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v35
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v70
									if v69 != 0 {
										v80 = int32(78674)
									} else {
										v80 = int32(78774)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v80
									F_errmsg_internal(m, int32(203982), v9)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										v88 = int32(927)
										F_errfinish(m, int32(493093), v88, int32(410716))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v94 != 0 {
												m.T0[v94].(func(*base.Module, int32))(m, l0)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
													v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													F_MemoryContextReset(m, v99)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return
													} else {
														m.G0 = v9 - int32(-64)
														return
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
												v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												F_MemoryContextReset(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													m.G0 = v9 - int32(-64)
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
	}
}
func F_tuplesort_putbrintuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(4515600)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v14
	v18 = F_palloc(m, l2+int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = l2
		if l2 != 0 {
			v23 = F__emscripten_memcpy_bulkmem(m, v18+int32(4), l1, l2)
			mBase = m.M
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v27 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(v27)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v26
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
		if v30&int32(2) != 0 {
			v33 = F_GetMemoryChunkSpace(m, v18)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v39 = v33
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v41 != 0 {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
					v45 = base.B2i32(v42 != int32(0))
				} else {
					v45 = int32(0)
				}
				F_tuplesort_puttuple_common(m, l0, v9, v45, v39)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
					m.G0 = v9 + int32(16)
					return
				}
			}
		} else {
			v39 = (l2 + int32(11)) & int32(-8)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v41 != 0 {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
				v45 = base.B2i32(v42 != int32(0))
			} else {
				v45 = int32(0)
			}
			F_tuplesort_puttuple_common(m, l0, v9, v45, v39)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F_tuplesort_putdatum(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v3 = l2
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(4515600)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v15
	if v3 == v4 {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v19 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)) = uint8(v24)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			v28 = F_datumCopy(m, l1, v24, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v28
				v31 = v28
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v31
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
				if v35 == int32(1) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
					v45 = (v3 ^ int32(1)) & base.B2i32(v41 != int32(0))
				} else {
					v45 = int32(0)
				}
				F_tuplesort_puttuple_common(m, l0, v10, v45, int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v13
					m.G0 = v10 + int32(16)
					return
				}
			}
		} else {
			v20 = l1
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)) = uint8(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v20
			v31 = v4
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v31
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
			if v35 == int32(1) {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
				v45 = (v3 ^ int32(1)) & base.B2i32(v41 != int32(0))
			} else {
				v45 = int32(0)
			}
			F_tuplesort_puttuple_common(m, l0, v10, v45, int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v13
				m.G0 = v10 + int32(16)
				return
			}
		}
	} else {
		v20 = v4
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)) = uint8(v3)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v20
		v31 = v4
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v31
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v35 == int32(1) {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
			v45 = (v3 ^ int32(1)) & base.B2i32(v41 != int32(0))
		} else {
			v45 = int32(0)
		}
		F_tuplesort_puttuple_common(m, l0, v10, v45, int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v13
			m.G0 = v10 + int32(16)
			return
		}
	}
}
func F_tuplesort_rescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	v3 = int32(4515600)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	switch v8 - int32(3) {
	case 0:
		v32 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v32)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v32
		v36 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+228)) = uint8(v36)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v36
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v4
		return
	case 1:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		F_LogicalTapeRewindForRead(m, v11, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = int64(0)
			v17 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v17)
			v36 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+228)) = uint8(v36)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v36
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v4
			return
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(352249), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errfinish(m, int32(493093), int32(2424), int32(284688))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_tuplesort_skiptuples(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if base.Ui32(int32(2)) <= base.Ui32(v10-int32(4)) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L31
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v92
L3:
	;
	v92 = int32(1)
	goto L2
L4:
	;
	if v10 != int32(3) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v49 = int32(4515600)
	v50 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v52
	if int64(0) < l1 {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if l1 <= base.I64_extend_i32_s(v17-v18) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v18 + base.I32_wrap_i64(l1)
	goto L3
L9:
	;
	goto L10
L10:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v25)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v17
	v28 = int32(0)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v29 != v25 {
		v92 = v28
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v17 < v32 {
		v92 = v28
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	F_errmsg_internal(m, int32(78838), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(493093), int32(1739), int32(164151))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v57 = l1
	goto L20
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v50
	goto L3
L20:
	;
	v62 = F_tuplesort_gettuple_common(m, l0, int32(1), v8)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L13
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	if v62 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v50
	v92 = int32(0)
	goto L2
L24:
	;
	goto L25
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v73 = int64(1)
	if base.Ui64(v73) < base.Ui64(v57) {
		v57 = v57 - v73
		goto L20
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	goto L21
L31:
	;
	F_errmsg_internal(m, int32(352249), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(493093), int32(1766), int32(164151))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
