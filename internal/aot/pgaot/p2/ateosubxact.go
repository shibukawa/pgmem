package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOSubXact_Inval(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Inval[0])) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Inval[1]))
	if v10 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Inval[2]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	goto L6
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	if v15 != v16 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if l0 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Inval[1])) = v95
	F_pfree(m, v10)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L12
	} else {
		goto L36
	}
L9:
	;
	F_CommandEndInvalidationMessages(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v53 < v54 {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	return
L13:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	if v20 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v29
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v38
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v43
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)))
	if v47 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v15-int32(1) <= v21 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v25 - int32(1)
	return
L18:
	;
	goto L17
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v51 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+16)) = uint8(v51)
	goto L21
L20:
	;
	goto L21
L21:
	;
	goto L8
L22:
	;
	v56 = v53
	goto L25
L23:
	;
	goto L24
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	if v72 < v73 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Inval[3]))
	F_LocalExecuteInvalidationMessage(m, v60+v56<<(uint(int32(4))%32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L27
	}
L26:
	;
	goto L24
L27:
	;
	v67 = v56 + int32(1)
	if v67 != v54 {
		v56 = v67
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v75 = v72
	goto L32
L30:
	;
	goto L31
L31:
	;
	goto L8
L32:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Inval[4]))
	F_LocalExecuteInvalidationMessage(m, v79+v75<<(uint(int32(4))%32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	v86 = v75 + int32(1)
	if v86 != v73 {
		v75 = v86
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L4
}
func F_AtEOSubXact_RelationCache(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_RelationCache[0])) = v4
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AtEOSubXact_RelationCache[1])))
	if v15 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_RelationCache[2]))
	if v19 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v47 = v9 + int32(12)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_RelationCache[3]))
	F_hash_seq_init(m, v47, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L15
	}
L5:
	;
	v25 = v4
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_RelationCache[3]))
	v34 = int32(0)
	v36 = F_hash_search(m, v29, v25<<(uint(int32(2))%32)+int32(_a_F_AtEOSubXact_RelationCache_0), v34, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	return
L9:
	;
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	F_AtEOSubXact_cleanup(m, v38, l0, l1, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v42 = v25 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_RelationCache[2]))
	if v42 < v44 {
		v25 = v42
		goto L6
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	goto L7
L15:
	;
	v52 = F_hash_seq_search(m, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if v52 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v59 = v52
	goto L18
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	F_AtEOSubXact_cleanup(m, v62, l0, l1, l2)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L20
	}
L19:
	;
	goto L1
L20:
	;
	v67 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	if v67 != 0 {
		v59 = v67
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
}
func F_AtEOSubXact_cleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 != l2 {
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v47 != l2 {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v59 != l2 {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v71 != l2 {
				} else {
					v74 = l0 + int32(44)
					if l1 != 0 {
						v84 = v74
						*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
					}
				}
			} else {
				if l1 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v78 != l2 {
					} else {
						v84 = l0 + int32(44)
						*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v71 != l2 {
					} else {
						v74 = l0 + int32(44)
						if l1 != 0 {
							v84 = v74
							*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
						}
					}
				}
			}
		} else {
			if l1 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v63 != l2 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
				}
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v71 != l2 {
				} else {
					v74 = l0 + int32(44)
					if l1 != 0 {
						v84 = v74
						*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = l3
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v59 != l2 {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v71 != l2 {
					} else {
						v74 = l0 + int32(44)
						if l1 != 0 {
							v84 = v74
							*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
						}
					}
				} else {
					if l1 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v78 != l2 {
						} else {
							v84 = l0 + int32(44)
							*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v71 != l2 {
						} else {
							v74 = l0 + int32(44)
							if l1 != 0 {
								v84 = v74
								*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
							}
						}
					}
				}
			}
		}
		m.G0 = v9 + int32(16)
		return
	} else {
		if l1 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v13 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l3
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v52 != l2 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = l3
				}
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v59 != l2 {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v71 != l2 {
					} else {
						v74 = l0 + int32(44)
						if l1 != 0 {
							v84 = v74
							*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
						}
					}
				} else {
					if l1 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v78 != l2 {
						} else {
							v84 = l0 + int32(44)
							*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v71 != l2 {
						} else {
							v74 = l0 + int32(44)
							if l1 != 0 {
								v84 = v74
								*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
							}
						}
					}
				}
				m.G0 = v9 + int32(16)
				return
			} else {
				v17 = l0 + int32(32)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v18 == int32(0) {
					v21 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v21
					*(*int64)(unsafe.Add(mBase, uint32(v17))) = v21
					F_RelationClearRelation(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = l3
					v30 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						if v30 == int32(0) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v47 != l2 {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v59 != l2 {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v71 != l2 {
									} else {
										v74 = l0 + int32(44)
										if l1 != 0 {
											v84 = v74
											*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
										}
									}
								} else {
									if l1 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v78 != l2 {
										} else {
											v84 = l0 + int32(44)
											*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v71 != l2 {
										} else {
											v74 = l0 + int32(44)
											if l1 != 0 {
												v84 = v74
												*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
											}
										}
									}
								}
							} else {
								if l1 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v63 != l2 {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
									}
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v71 != l2 {
									} else {
										v74 = l0 + int32(44)
										if l1 != 0 {
											v84 = v74
											*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = l3
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v59 != l2 {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v71 != l2 {
										} else {
											v74 = l0 + int32(44)
											if l1 != 0 {
												v84 = v74
												*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
											}
										}
									} else {
										if l1 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v78 != l2 {
											} else {
												v84 = l0 + int32(44)
												*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v71 != l2 {
											} else {
												v74 = l0 + int32(44)
												if l1 != 0 {
													v84 = v74
													*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
												}
											}
										}
									}
								}
							}
							m.G0 = v9 + int32(16)
							return
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v34 + int32(4)
							F_errmsg_internal(m, int32(_a_F_AtEOSubXact_cleanup_0), v9)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_AtEOSubXact_cleanup_1), int32(3475), int32(_a_F_AtEOSubXact_cleanup_2))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v47 != l2 {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v59 != l2 {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v71 != l2 {
											} else {
												v74 = l0 + int32(44)
												if l1 != 0 {
													v84 = v74
													*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
												}
											}
										} else {
											if l1 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v78 != l2 {
												} else {
													v84 = l0 + int32(44)
													*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v71 != l2 {
												} else {
													v74 = l0 + int32(44)
													if l1 != 0 {
														v84 = v74
														*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
													}
												}
											}
										}
									} else {
										if l1 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v63 != l2 {
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
											}
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v71 != l2 {
											} else {
												v74 = l0 + int32(44)
												if l1 != 0 {
													v84 = v74
													*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = l3
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v59 != l2 {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v71 != l2 {
												} else {
													v74 = l0 + int32(44)
													if l1 != 0 {
														v84 = v74
														*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
													}
												}
											} else {
												if l1 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
													v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v78 != l2 {
													} else {
														v84 = l0 + int32(44)
														*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
													v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v71 != l2 {
													} else {
														v74 = l0 + int32(44)
														if l1 != 0 {
															v84 = v74
															*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
														}
													}
												}
											}
										}
									}
									m.G0 = v9 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v17 = l0 + int32(32)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v18 == int32(0) {
				v21 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v21
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = v21
				F_RelationClearRelation(m, l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = l3
				v30 = F_errstart(m, int32(19), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					if v30 == int32(0) {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v47 != l2 {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v59 != l2 {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v71 != l2 {
								} else {
									v74 = l0 + int32(44)
									if l1 != 0 {
										v84 = v74
										*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
									}
								}
							} else {
								if l1 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v78 != l2 {
									} else {
										v84 = l0 + int32(44)
										*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v71 != l2 {
									} else {
										v74 = l0 + int32(44)
										if l1 != 0 {
											v84 = v74
											*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
										}
									}
								}
							}
						} else {
							if l1 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v63 != l2 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
								}
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v71 != l2 {
								} else {
									v74 = l0 + int32(44)
									if l1 != 0 {
										v84 = v74
										*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = l3
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v59 != l2 {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v71 != l2 {
									} else {
										v74 = l0 + int32(44)
										if l1 != 0 {
											v84 = v74
											*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
										}
									}
								} else {
									if l1 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v78 != l2 {
										} else {
											v84 = l0 + int32(44)
											*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v71 != l2 {
										} else {
											v74 = l0 + int32(44)
											if l1 != 0 {
												v84 = v74
												*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
											}
										}
									}
								}
							}
						}
						m.G0 = v9 + int32(16)
						return
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v34 + int32(4)
						F_errmsg_internal(m, int32(_a_F_AtEOSubXact_cleanup_0), v9)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_AtEOSubXact_cleanup_1), int32(3475), int32(_a_F_AtEOSubXact_cleanup_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v47 != l2 {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v59 != l2 {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v71 != l2 {
										} else {
											v74 = l0 + int32(44)
											if l1 != 0 {
												v84 = v74
												*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
											}
										}
									} else {
										if l1 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v78 != l2 {
											} else {
												v84 = l0 + int32(44)
												*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v71 != l2 {
											} else {
												v74 = l0 + int32(44)
												if l1 != 0 {
													v84 = v74
													*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
												}
											}
										}
									}
								} else {
									if l1 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v63 != l2 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
										}
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v71 != l2 {
										} else {
											v74 = l0 + int32(44)
											if l1 != 0 {
												v84 = v74
												*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = l3
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v59 != l2 {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v71 != l2 {
											} else {
												v74 = l0 + int32(44)
												if l1 != 0 {
													v84 = v74
													*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
												}
											}
										} else {
											if l1 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l3
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v78 != l2 {
												} else {
													v84 = l0 + int32(44)
													*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)))) = int32(0)
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v71 != l2 {
												} else {
													v74 = l0 + int32(44)
													if l1 != 0 {
														v84 = v74
														*(*int32)(unsafe.Add(mBase, uint32(v84))) = l3
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
													}
												}
											}
										}
									}
								}
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
