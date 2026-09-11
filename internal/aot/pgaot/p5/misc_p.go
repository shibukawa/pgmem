package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParamsErrorCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	if l0 == int32(0) {
		m.G0 = v6 + int32(32)
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 == int32(0) {
			m.G0 = v6 + int32(32)
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
			if v13 == int32(0) {
				m.G0 = v6 + int32(32)
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v16 == int32(0) {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
						F_errcontext_msg(m, int32(_a_F_ParamsErrorCallback_0), v6)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							m.G0 = v6 + int32(32)
							return
						}
					}
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if v19 == int32(0) {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
							F_errcontext_msg(m, int32(_a_F_ParamsErrorCallback_0), v6)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								m.G0 = v6 + int32(32)
								return
							}
						}
					} else {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v27
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v25
							F_errcontext_msg(m, int32(_a_F_ParamsErrorCallback_1), v6+int32(16))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								m.G0 = v6 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_PartConstraintImpliedByRelConstraint(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v11 == v3 {
		v82 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v87 = F_ConstraintImpliedByRelConstraint(m, l0, l1, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L10
	} else {
		goto L15
	}
L2:
	;
	v14 = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v15 != v14 {
		v82 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v18 <= int32(0) {
		v82 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = v3
	v26 = v14
	goto L5
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = v26 - int32(1)
	v35 = v30 + v32<<(uint(int32(4))%32)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+31)))
	if v36 != int32(118) {
		v73 = v25
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v82 = v73
	goto L1
L7:
	;
	v76 = v26 + int32(1)
	if v76 <= v18 {
		v25 = v73
		v26 = v76
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+29)))
	if v39 != 0 {
		v73 = v25
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v42 = F_palloc0(m, int32(20))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(52)
	v55 = v30 + v40<<(uint(int32(4))%32) + v32*int32(100)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+88))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+96))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+116))
	v60 = F_makeVar(m, int32(1), base.I32_extend16_s(v26), v56, v57, v58, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = int32(-1)
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+12)) = uint8(v64)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v60
	v69 = F_lappend(m, v25, v42)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v73 = v69
	goto L7
L14:
	;
	goto L6
L15:
	;
	return v87
}
func F_PhysicalReplicationSlotNewXmin(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_PhysicalReplicationSlotNewXmin[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	if v8 != 0 {
		F_s_lock(m, v7, int32(_a_F_PhysicalReplicationSlotNewXmin_0), int32(2527), int32(_a_F_PhysicalReplicationSlotNewXmin_1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_PhysicalReplicationSlotNewXmin[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = int32(0)
			if base.Ui32(l0) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
				v45 = int32(1)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+96))
				if base.Ui32(v22) < base.Ui32(int32(3)) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
					v45 = int32(1)
				} else {
					v25 = int32(0)
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v22)) == v25 {
						v37 = base.B2i32(base.Ui32(v22) < base.Ui32(l0))
					} else {
						v37 = int32(base.Ui32(v22-l0) >> (uint(int32(31)) % 32))
					}
					if v37 == int32(0) {
						v45 = v25
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
						v45 = int32(1)
					}
				}
			}
			if base.Ui32(l1) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
				F_ReplicationSlotMarkDirty(m)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+100))
				if base.Ui32(v49) < base.Ui32(int32(3)) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
					F_ReplicationSlotMarkDirty(m)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v49)) == int32(0) {
						v63 = base.B2i32(base.Ui32(v49) < base.Ui32(l1))
					} else {
						v63 = int32(base.Ui32(v49-l1) >> (uint(int32(31)) % 32))
					}
					if v63 == int32(0) {
						v71 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v71
						if v45 == v71 {
							return
						} else {
							F_ReplicationSlotMarkDirty(m)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
						F_ReplicationSlotMarkDirty(m)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_PhysicalReplicationSlotNewXmin[1]))
		*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = int32(0)
		if base.Ui32(l0) < base.Ui32(int32(3)) {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
			v45 = int32(1)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+96))
			if base.Ui32(v22) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
				v45 = int32(1)
			} else {
				v25 = int32(0)
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v22)) == v25 {
					v37 = base.B2i32(base.Ui32(v22) < base.Ui32(l0))
				} else {
					v37 = int32(base.Ui32(v22-l0) >> (uint(int32(31)) % 32))
				}
				if v37 == int32(0) {
					v45 = v25
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
					v45 = int32(1)
				}
			}
		}
		if base.Ui32(l1) < base.Ui32(int32(3)) {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
			F_ReplicationSlotMarkDirty(m)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+100))
			if base.Ui32(v49) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
				F_ReplicationSlotMarkDirty(m)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v49)) == int32(0) {
					v63 = base.B2i32(base.Ui32(v49) < base.Ui32(l1))
				} else {
					v63 = int32(base.Ui32(v49-l1) >> (uint(int32(31)) % 32))
				}
				if v63 == int32(0) {
					v71 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v71
					if v45 == v71 {
						return
					} else {
						F_ReplicationSlotMarkDirty(m)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
					F_ReplicationSlotMarkDirty(m)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_PopActiveSnapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v8 - int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v13 != 0 {
		v19 = v5
		F_pfree(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0])) = v6
			if v6 != 0 {
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
				if v26 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
					v33 = v31 - int32(48)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v34))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v29)) == int32(0) {
						v46 = base.B2i32(base.Ui32(v29) < base.Ui32(v34))
					} else {
						v46 = int32(base.Ui32(v29-v34) >> (uint(int32(31)) % 32))
					}
					if v46 == int32(0) {
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v50 = v49
						*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
						v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
					}
				} else {
					v50 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
					v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
				}
			}
			return
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
		if v14 != 0 {
			v19 = v5
			F_pfree(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0])) = v6
				if v6 != 0 {
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
					if v26 != 0 {
						v28 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
						v33 = v31 - int32(48)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v34))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v29)) == int32(0) {
							v46 = base.B2i32(base.Ui32(v29) < base.Ui32(v34))
						} else {
							v46 = int32(base.Ui32(v29-v34) >> (uint(int32(31)) % 32))
						}
						if v46 == int32(0) {
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							v50 = v49
							*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
						}
					} else {
						v50 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
						v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
					}
				}
				return
			}
		} else {
			F_pfree(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0]))
				v19 = v18
				F_pfree(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0])) = v6
					if v6 != 0 {
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
						if v26 != 0 {
							v28 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
							v31 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
							v33 = v31 - int32(48)
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v34))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v29)) == int32(0) {
								v46 = base.B2i32(base.Ui32(v29) < base.Ui32(v34))
							} else {
								v46 = int32(base.Ui32(v29-v34) >> (uint(int32(31)) % 32))
							}
							if v46 == int32(0) {
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
								v50 = v49
								*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
								v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
							}
						} else {
							v50 = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
						}
					}
					return
				}
			}
		}
	}
}
func F_PostgresSendReadyForQueryIfNecessary(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])))
	if v4 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[1]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		if (v9-int32(7))&int32(-9) == int32(0) {
			v17 = int32(0)
			F_pgstat_report_activity(m, int32(6), v17)
			mBase = m.M
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[2]))
			if v20 <= v17 {
				F_ReportChangedGUCOptions(m)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
					F_ReadyForQuery(m, v99)
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return
					} else {
						v103 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
						return
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[4]))
				if v24 != 0 {
					v27 = base.B2i32(v24 <= v20)
				} else {
					v27 = int32(0)
				}
				if v27 != 0 {
					F_ReportChangedGUCOptions(m)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
						F_ReadyForQuery(m, v99)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							v103 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
							return
						}
					}
				} else {
					v29 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[5])) = uint8(v29)
					F_enable_timeout_after(m, int32(7), v20)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_ReportChangedGUCOptions(m)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
							F_ReadyForQuery(m, v99)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								v103 = int32(0)
								*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
								return
							}
						}
					}
				}
			}
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[1]))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
			if v36 != int32(0) {
				v40 = int32(0)
				F_pgstat_report_activity(m, int32(4), v40)
				mBase = m.M
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[2]))
				if v43 <= v40 {
					F_ReportChangedGUCOptions(m)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
						F_ReadyForQuery(m, v99)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							v103 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
							return
						}
					}
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[4]))
					if v47 != 0 {
						v50 = base.B2i32(v47 <= v43)
					} else {
						v50 = int32(0)
					}
					if v50 != 0 {
						F_ReportChangedGUCOptions(m)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
							F_ReadyForQuery(m, v99)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								v103 = int32(0)
								*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
								return
							}
						}
					} else {
						v52 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[5])) = uint8(v52)
						F_enable_timeout_after(m, int32(7), v43)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_ReportChangedGUCOptions(m)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
								F_ReadyForQuery(m, v99)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									v103 = int32(0)
									*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
									return
								}
							}
						}
					}
				}
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[6]))
				if v58 != 0 {
					F_ProcessNotifyInterrupt(m, int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v63 = F_pgstat_report_stat(m, int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[7])))
							if int32(0) < v63 {
								if v70 != 0 {
									v82 = int32(0)
									F_pgstat_report_activity(m, int32(2), v82)
									mBase = m.M
									v85 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
									if v85 <= v82 {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
												return
											}
										}
									} else {
										v89 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v89)
										F_enable_timeout_after(m, int32(9), v85)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
													return
												}
											}
										}
									}
								} else {
									F_enable_timeout_after(m, int32(10), v63)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										v82 = int32(0)
										F_pgstat_report_activity(m, int32(2), v82)
										mBase = m.M
										v85 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
										if v85 <= v82 {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
													return
												}
											}
										} else {
											v89 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v89)
											F_enable_timeout_after(m, int32(9), v85)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												F_ReportChangedGUCOptions(m)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
													F_ReadyForQuery(m, v99)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return
													} else {
														v103 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
														return
													}
												}
											}
										}
									}
								}
							} else {
								if v70 == int32(0) {
									v82 = int32(0)
									F_pgstat_report_activity(m, int32(2), v82)
									mBase = m.M
									v85 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
									if v85 <= v82 {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
												return
											}
										}
									} else {
										v89 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v89)
										F_enable_timeout_after(m, int32(9), v85)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
													return
												}
											}
										}
									}
								} else {
									F_disable_timeout(m, int32(10))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										v82 = int32(0)
										F_pgstat_report_activity(m, int32(2), v82)
										mBase = m.M
										v85 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
										if v85 <= v82 {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
													return
												}
											}
										} else {
											v89 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v89)
											F_enable_timeout_after(m, int32(9), v85)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												F_ReportChangedGUCOptions(m)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
													F_ReadyForQuery(m, v99)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return
													} else {
														v103 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
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
					v63 = F_pgstat_report_stat(m, int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[7])))
						if int32(0) < v63 {
							if v70 != 0 {
								v82 = int32(0)
								F_pgstat_report_activity(m, int32(2), v82)
								mBase = m.M
								v85 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
								if v85 <= v82 {
									F_ReportChangedGUCOptions(m)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
										F_ReadyForQuery(m, v99)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											v103 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
											return
										}
									}
								} else {
									v89 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v89)
									F_enable_timeout_after(m, int32(9), v85)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
												return
											}
										}
									}
								}
							} else {
								F_enable_timeout_after(m, int32(10), v63)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v82 = int32(0)
									F_pgstat_report_activity(m, int32(2), v82)
									mBase = m.M
									v85 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
									if v85 <= v82 {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
												return
											}
										}
									} else {
										v89 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v89)
										F_enable_timeout_after(m, int32(9), v85)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
													return
												}
											}
										}
									}
								}
							}
						} else {
							if v70 == int32(0) {
								v82 = int32(0)
								F_pgstat_report_activity(m, int32(2), v82)
								mBase = m.M
								v85 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
								if v85 <= v82 {
									F_ReportChangedGUCOptions(m)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
										F_ReadyForQuery(m, v99)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											v103 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
											return
										}
									}
								} else {
									v89 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v89)
									F_enable_timeout_after(m, int32(9), v85)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
												return
											}
										}
									}
								}
							} else {
								F_disable_timeout(m, int32(10))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									v82 = int32(0)
									F_pgstat_report_activity(m, int32(2), v82)
									mBase = m.M
									v85 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
									if v85 <= v82 {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v99)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
												return
											}
										}
									} else {
										v89 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v89)
										F_enable_timeout_after(m, int32(9), v85)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v103)
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
		return
	}
}
func F_PreventCommandDuringRecovery(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PreventCommandDuringRecovery[0])))
	if v9 == int32(1) {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_PreventCommandDuringRecovery[1]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+316))
		v17 = base.B2i32(v15 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_PreventCommandDuringRecovery[0])) = uint8(v17)
		v19 = v17
	} else {
		v19 = int32(0)
	}
	if v19 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(100663618))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(_a_F_PreventCommandDuringRecovery_0), v5)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PreventCommandDuringRecovery_1), int32(448), int32(_a_F_PreventCommandDuringRecovery_2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_ProcessCommittedInvalidationMessages(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l1 <= int32(0) {
		m.G0 = v9 + int32(32)
		return
	} else {
		v15 = F_errstart(m, int32(11), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
				if l2 != 0 {
					v20 = int32(_a_F_ProcessCommittedInvalidationMessages_0)
				} else {
					v20 = int32(_a_F_ProcessCommittedInvalidationMessages_1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v20
				F_errmsg_internal(m, int32(_a_F_ProcessCommittedInvalidationMessages_2), v9+int32(16))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ProcessCommittedInvalidationMessages_3), int32(1143), int32(_a_F_ProcessCommittedInvalidationMessages_4))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						if l2 != 0 {
							v34 = F_errstart(m, int32(11), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								if v34 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
									F_errmsg_internal(m, int32(_a_F_ProcessCommittedInvalidationMessages_5), v9)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ProcessCommittedInvalidationMessages_3), int32(1147), int32(_a_F_ProcessCommittedInvalidationMessages_4))
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return
										} else {
											if l3 != 0 {
												v46 = F_GetDatabasePath(m, l3, l4)
												mBase = m.M
												v47 = m.ExcPending
												if v47 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
													F_RelationCacheInitFilePreInvalidate(m)
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return
													} else {
														v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
														F_pfree(m, v52)
														mBase = m.M
														v54 = m.ExcPending
														if v54 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
															F_SendSharedInvalidMessages(m, l0, l1)
															mBase = m.M
															v61 = m.ExcPending
															if v61 != 0 {
																return
															} else {
																F_RelationCacheInitFilePostInvalidate(m)
																mBase = m.M
																v63 = m.ExcPending
																if v63 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													}
												}
											} else {
												F_RelationCacheInitFilePreInvalidate(m)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													F_SendSharedInvalidMessages(m, l0, l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_RelationCacheInitFilePostInvalidate(m)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								} else {
									if l3 != 0 {
										v46 = F_GetDatabasePath(m, l3, l4)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
											F_RelationCacheInitFilePreInvalidate(m)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
												F_pfree(m, v52)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
													F_SendSharedInvalidMessages(m, l0, l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_RelationCacheInitFilePostInvalidate(m)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									} else {
										F_RelationCacheInitFilePreInvalidate(m)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_SendSharedInvalidMessages(m, l0, l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_RelationCacheInitFilePostInvalidate(m)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						} else {
							F_SendSharedInvalidMessages(m, l0, l1)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			} else {
				if l2 != 0 {
					v34 = F_errstart(m, int32(11), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						if v34 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
							F_errmsg_internal(m, int32(_a_F_ProcessCommittedInvalidationMessages_5), v9)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ProcessCommittedInvalidationMessages_3), int32(1147), int32(_a_F_ProcessCommittedInvalidationMessages_4))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									if l3 != 0 {
										v46 = F_GetDatabasePath(m, l3, l4)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
											F_RelationCacheInitFilePreInvalidate(m)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
												F_pfree(m, v52)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
													F_SendSharedInvalidMessages(m, l0, l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_RelationCacheInitFilePostInvalidate(m)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									} else {
										F_RelationCacheInitFilePreInvalidate(m)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_SendSharedInvalidMessages(m, l0, l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_RelationCacheInitFilePostInvalidate(m)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						} else {
							if l3 != 0 {
								v46 = F_GetDatabasePath(m, l3, l4)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
									F_RelationCacheInitFilePreInvalidate(m)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
										F_pfree(m, v52)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
											F_SendSharedInvalidMessages(m, l0, l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_RelationCacheInitFilePostInvalidate(m)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							} else {
								F_RelationCacheInitFilePreInvalidate(m)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_SendSharedInvalidMessages(m, l0, l1)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										F_RelationCacheInitFilePostInvalidate(m)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_SendSharedInvalidMessages(m, l0, l1)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_ProcessInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v319 int64
	_ = v319
	var v325 int64
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int64
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v769 int32
	_ = v769
	var v770 int64
	_ = v770
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	if v16 != 0 {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L18
	} else {
		goto L361
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L18
	} else {
		goto L357
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L18
	} else {
		goto L353
	}
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L18
	} else {
		goto L349
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1])) = int32(0)
	F_LockErrorCleanup(m)
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L18
	} else {
		goto L344
	}
L6:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L18
	} else {
		goto L339
	}
L7:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L18
	} else {
		goto L334
	}
L8:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L18
	} else {
		goto L329
	}
L9:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L18
	} else {
		goto L325
	}
L10:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L18
	} else {
		goto L321
	}
L11:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L18
	} else {
		goto L317
	}
L12:
	;
	m.G0 = v13 + int32(32)
	return
L13:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2]))
	if v18 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[3])) = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[4]))
	if v23 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[4])) = v25
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1])) = v25
	F_LockErrorCleanup(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[5]))
	if v148 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L18:
	;
	return
L19:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[6])))
	if v33 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v33 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[7]))
	if v37 != int32(2) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[7])) = int32(0)
	goto L20
L23:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[8]))
	if v44 == int32(4) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[9]))
	if v48 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[10]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[11]))
	if v51 == v53 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v57 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L18
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[8]))
	if v72 != int32(12) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	if v57 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_0), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L18
	} else {
		goto L35
	}
L33:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3435), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L18
	} else {
		goto L57
	}
L37:
	;
	switch v72 - int32(5) {
	case 0:
		goto L40
	default:
		goto L36
	case 9:
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v117 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L18
	} else {
		goto L50
	}
L40:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L18
	} else {
		goto L46
	}
L41:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_3), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3446), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v101 + int32(96)
	F_errmsg(m, int32(_a_F_ProcessInterrupts_4), v13+int32(16))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3451), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	if v117 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_5), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L18
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L18
	} else {
		goto L56
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3455), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L18
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_6), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L18
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3462), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L18
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[13]))
	if v265 != 0 {
		goto L5
	} else {
		goto L85
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[5])) = int32(0)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[14])))
	if v155 != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[15]))
	if v157 <= int32(0) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v160 = m.G0
	v162 = v160 - int32(48)
	m.G0 = v162
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[16]))
	v166 = int32(0)
	F_ModifyWaitEvent(m, v165, v166, int32(128), v166)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[16]))
	v174 = int32(0)
	v177 = F_WaitEventSetWait(m, v173, v174, v162, int32(3), v174)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L18
	} else {
		goto L67
	}
L66:
	;
	m.G0 = v162 + int32(48)
	if v231 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L67:
	;
	if v177 <= int32(0) {
		v231 = int32(1)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v184 = v177
	goto L69
L69:
	;
	v193 = int32(0)
	goto L71
L70:
	;
	v231 = int32(1)
	goto L66
L71:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v162+v193<<(uint(int32(4))%32))+4))
	v207 = v205 & int32(128)
	v209 = base.B2i32(v207 == int32(0))
	if v207 != 0 {
		v231 = v209
		goto L66
	} else {
		goto L73
	}
L72:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = int32(0)
	goto L78
L73:
	;
	if v205&int32(1) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v215 = v193 + int32(1)
	if v215 == v184 {
		v231 = v209
		goto L66
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	goto L72
L77:
	;
	v193 = v215
	goto L71
L78:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[16]))
	v224 = int32(0)
	v227 = F_WaitEventSetWait(m, v223, v224, v162, int32(3), v224)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L18
	} else {
		goto L79
	}
L79:
	;
	if int32(0) < v227 {
		v184 = v227
		goto L69
	} else {
		goto L80
	}
L80:
	;
	goto L70
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[13])) = int32(1)
	goto L61
L82:
	;
	goto L83
L83:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[15]))
	F_enable_timeout_after(m, int32(11), v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	goto L61
L85:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1]))
	if v267 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[18]))
	if v338 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L87:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1]))
	if v278 == int32(0) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[19]))
	if v271 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[3])) = int32(1)
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1])) = int32(0)
	v289 = int32(_a_F_ProcessInterrupts_7)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[20])))
	if v290&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v302 = int32(_a_F_ProcessInterrupts_8)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[21])))
	if v303&int32(1) != 0 {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v293 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[20])) = uint8(v293)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v296 = v290 & int32(1)
	goto L91
L95:
	;
	if v296 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[21])) = uint8(v306)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v309 = v303 & int32(1)
	goto L95
L99:
	;
	if v309 != 0 {
		goto L6
	} else {
		goto L107
	}
L100:
	;
	if v296 != 0 {
		goto L1
	} else {
		goto L106
	}
L101:
	;
	if v309 == int32(0) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v319 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessInterrupts[22]))
	goto L103
L103:
	;
	v325 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessInterrupts[23]))
	goto L104
L104:
	;
	if v319 < v325 {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	goto L1
L106:
	;
	goto L99
L107:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[8]))
	if v328 == int32(4) {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[14])))
	if v332 == int32(0) {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	goto L86
L110:
	;
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[24]))
	if v500 != 0 {
		goto L163
	} else {
		goto L164
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[18])) = int32(0)
	v345 = int32(7)
	goto L112
L112:
	;
	v356 = v345 << (uint(int32(2)) % 32)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v356)+uint32(_c_F_ProcessInterrupts[25])))
	if v359 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L110
L114:
	;
	v486 = v345 + int32(1)
	if v486 != int32(14) {
		v345 = v486
		goto L112
	} else {
		goto L162
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356)+uint32(_c_F_ProcessInterrupts[25]))) = int32(0)
	switch v345 - int32(7) {
	case 0:
		goto L116
	case 1, 2, 3:
		goto L118
	case 4:
		goto L117
	case 5:
		goto L120
	case 6:
		goto L122
	default:
		goto L121
	}
L116:
	;
	F_pgstat_report_recovery_conflict(m, v345)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L18
	} else {
		goto L152
	}
L117:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26]))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+24))
	goto L139
L118:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26]))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+24))
	goto L134
L119:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[27]))
	v398 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+73)) = uint8(v398)
	goto L118
L120:
	;
	v392 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L18
	} else {
		goto L132
	}
L121:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L18
	} else {
		goto L129
	}
L122:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[28]))
	if v367 == int32(0) {
		goto L110
	} else {
		goto L123
	}
L123:
	;
	v370 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L18
	} else {
		goto L124
	}
L124:
	;
	if v370 != 0 {
		goto L119
	} else {
		goto L125
	}
L125:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[29]))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+72))
	goto L126
L126:
	;
	if int32(0) <= v374 {
		goto L110
	} else {
		goto L127
	}
L127:
	;
	F_CheckDeadLockAlert(m)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L18
	} else {
		goto L128
	}
L128:
	;
	goto L110
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v345
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_9), v13)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L18
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3358), int32(_a_F_ProcessInterrupts_10))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L18
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	if v392 == int32(0) {
		goto L114
	} else {
		goto L133
	}
L133:
	;
	goto L119
L134:
	;
	if base.B2i32(v402 != int32(0)) == int32(0) {
		goto L114
	} else {
		goto L135
	}
L135:
	;
	if v345 == int32(11) {
		goto L117
	} else {
		goto L136
	}
L136:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26]))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+28))
	goto L137
L137:
	;
	if int32(1) < v411 {
		goto L116
	} else {
		goto L138
	}
L138:
	;
	goto L117
L139:
	;
	if (v416-int32(7))&int32(-9) == int32(0) {
		goto L114
	} else {
		goto L140
	}
L140:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[14])))
	if v424 != 0 {
		goto L116
	} else {
		goto L141
	}
L141:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[19]))
	if v426 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v427 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v356)+uint32(_c_F_ProcessInterrupts[25]))) = v427
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[18])) = v427
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[3])) = v427
	goto L114
L143:
	;
	goto L144
L144:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	F_pgstat_report_recovery_conflict(m, v345)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L18
	} else {
		goto L147
	}
L147:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L18
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_11), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errdetail_recovery_conflict(m, v345)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3331), int32(_a_F_ProcessInterrupts_10))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L18
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	if v345 == int32(7) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v467 = int32(67240389)
	goto L156
L155:
	;
	v467 = int32(16777220)
	goto L156
L156:
	;
	F_errcode(m, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_12), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L18
	} else {
		goto L158
	}
L158:
	;
	F_errdetail_recovery_conflict(m, v345)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L18
	} else {
		goto L159
	}
L159:
	;
	F_errhint(m, int32(_a_F_ProcessInterrupts_13), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L18
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3354), int32(_a_F_ProcessInterrupts_10))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L18
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	goto L113
L163:
	;
	v502 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[24])) = v502
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[30]))
	if v502 < v505 {
		goto L11
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[31]))
	if v509 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L165
L167:
	;
	v511 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[31])) = v511
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[32]))
	if v511 < v514 {
		goto L10
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[33]))
	if v518 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	goto L169
L171:
	;
	v520 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[33])) = v520
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[34]))
	if v520 < v523 {
		goto L9
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[35]))
	if v527 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L173
L175:
	;
	v546 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[36]))
	if v546 != 0 {
		goto L181
	} else {
		goto L182
	}
L176:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[14])))
	if v531 == int32(0) {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26]))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+24))
	goto L178
L178:
	;
	if v536 != int32(0) {
		goto L175
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[35])) = int32(0)
	v543 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L18
	} else {
		goto L180
	}
L180:
	;
	goto L175
L181:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L18
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v550 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[37]))
	if v550 != 0 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	goto L183
L185:
	;
	v551 = m.G0
	v553 = v551 - int32(160)
	m.G0 = v553
	v555 = int32(_a_F_ProcessInterrupts_14)
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v557 + int32(1)
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[38]))
	if v562 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	goto L187
L187:
	;
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[39]))
	if v900 != 0 {
		goto L264
	} else {
		goto L265
	}
L188:
	;
	v580 = int32(_a_F_ProcessInterrupts_15)
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40])) = v579
	v585 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[37])) = v585
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41]))
	if v588 == v585 {
		v867 = v579
		goto L194
	} else {
		goto L195
	}
L189:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[42]))
	v572 = F_AllocSetContextCreateInternal(m, v567, int32(_a_F_ProcessInterrupts_16), int32(0), int32(_a_F_ProcessInterrupts_17), int32(_a_F_ProcessInterrupts_18))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L18
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	F_MemoryContextReset(m, v562)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L18
	} else {
		goto L193
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[38])) = v572
	v579 = v572
	goto L188
L193:
	;
	v578 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[38]))
	v579 = v578
	goto L188
L194:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40])) = v581
	F_MemoryContextReset(m, v867)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L18
	} else {
		goto L263
	}
L195:
	;
	if v588 == int32(_a_F_ProcessInterrupts_19) {
		v867 = v579
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v595 = v588
	goto L197
L197:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v595)+56))
	if v603 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v865 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[38]))
	v867 = v865
	goto L194
L199:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v861 != int32(_a_F_ProcessInterrupts_19) {
		v595 = v861
		goto L197
	} else {
		goto L262
	}
L200:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v595)+20))
	if v606 <= int32(0) {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v613 = int32(0)
	goto L202
L202:
	;
	v621 = v613 << (uint(int32(3)) % 32)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v595)+56))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v621+v622)+4))
	if v624 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	goto L199
L204:
	;
	v848 = v613 + int32(1)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v595)+20))
	if v848 < v849 {
		v613 = v848
		goto L202
	} else {
		goto L261
	}
L205:
	;
	v628 = v624
	goto L206
L206:
	;
	v642 = F_shm_mq_receive(m, v628, v553+int32(56), v553+int32(52), int32(1))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L18
	} else {
		goto L208
	}
L207:
	;
	goto L204
L208:
	;
	if v642 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	if v642 == int32(1) {
		goto L204
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	F_initStringInfo(m, v553+int32(36))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L18
	} else {
		goto L217
	}
L212:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L18
	} else {
		goto L213
	}
L213:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L18
	} else {
		goto L214
	}
L214:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_20), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L18
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_21), int32(1127), int32(_a_F_ProcessInterrupts_16))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L18
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v553)+52))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v553)+56))
	F_appendBinaryStringInfo(m, v553+int32(36), v668, v669)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L18
	} else {
		goto L218
	}
L218:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v595)+64))
	if v672 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v686 = F_pq_getmsgbyte(m, v553+int32(36))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L18
	} else {
		goto L228
	}
L220:
	;
	v675 = v672 + v613
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	if v676 != 0 {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v677 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v675))) = uint8(v677)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v595)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v595)+60)) = v679 + v677
	goto L219
L222:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v553)+36))
	F_pfree(m, v831)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L18
	} else {
		goto L259
	}
L223:
	;
	v811 = F_pq_getmsgint(m, v553+int32(36), int32(4))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L18
	} else {
		goto L254
	}
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L18
	} else {
		goto L251
	}
L225:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v595)+56))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v784+v621)+4))
	F_shm_mq_detach(m, v786)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L18
	} else {
		goto L250
	}
L226:
	;
	v736 = F_pq_getmsgint(m, v553+int32(36), int32(4))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L18
	} else {
		goto L243
	}
L227:
	;
	F_pq_parse_errornotice(m, v553+int32(36), v553+int32(60))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L18
	} else {
		goto L229
	}
L228:
	;
	v688 = base.I32_extend8_s(v686)
	switch v688 - int32(65) {
	case 0:
		goto L223
	default:
		goto L224
	case 4, 13:
		goto L227
	case 15:
		goto L226
	case 23:
		goto L225
	}
L229:
	;
	v697 = int32(21)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v553)+60))
	if v697 <= v698 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v701 = v697
	goto L232
L231:
	;
	v701 = v698
	goto L232
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v553)+60)) = v701
	v704 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[43]))
	if v704 != int32(2) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v553)+108))
	if v707 != 0 {
		goto L237
	} else {
		goto L238
	}
L234:
	;
	goto L235
L235:
	;
	v722 = int32(_a_F_ProcessInterrupts_22)
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[44]))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v595)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[44])) = v725
	F_ThrowErrorData(m, v553+int32(60))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L18
	} else {
		goto L242
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v553)+108)) = v719
	goto L235
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v553)+20)) = int32(_a_F_ProcessInterrupts_23)
	*(*int32)(unsafe.Add(mBase, uint32(v553)+16)) = v707
	v714 = F_psprintf(m, int32(_a_F_ProcessInterrupts_24), v553+int32(16))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L18
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v717 = F_pstrdup(m, int32(_a_F_ProcessInterrupts_23))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L18
	} else {
		goto L241
	}
L240:
	;
	v719 = v714
	goto L236
L241:
	;
	v719 = v717
	goto L236
L242:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[44])) = v723
	goto L222
L243:
	;
	v740 = F_pq_getmsgint64(m, v553+int32(36))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L18
	} else {
		goto L244
	}
L244:
	;
	F_pq_getmsgend(m, v553+int32(36))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L18
	} else {
		goto L245
	}
L245:
	;
	v748 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[45]))
	if v748 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	goto L222
L247:
	;
	goto L246
L248:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[46])))
	if v752 != int32(1) {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v755 = int32(_a_F_ProcessInterrupts_25)
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2]))
	v758 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2])) = v757 + v758
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v748)))
	*(*int32)(unsafe.Add(mBase, uint32(v748))) = v761 + v758
	v769 = v748 + v736<<(uint(int32(3))%32) + int32(232)
	v770 = *(*int64)(unsafe.Add(mBase, uint32(v769)))
	*(*int64)(unsafe.Add(mBase, uint32(v769))) = v770 + v740
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v748)))
	*(*int32)(unsafe.Add(mBase, uint32(v748))) = v773 + v758
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2])) = v779 - v758
	goto L247
L250:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v595)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v789+v621)+4)) = int32(0)
	goto L222
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v553))) = v688
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v553)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v553)+4)) = v798
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_26), v553)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L18
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_21), int32(1249), int32(_a_F_ProcessInterrupts_27))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L18
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	v815 = F_pq_getmsgrawstring(m, v553+int32(36))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L18
	} else {
		goto L255
	}
L255:
	;
	v819 = F_pq_getmsgrawstring(m, v553+int32(36))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L18
	} else {
		goto L256
	}
L256:
	;
	F_pq_endmessage(m, v553+int32(36))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L18
	} else {
		goto L257
	}
L257:
	;
	F_NotifyMyFrontEnd(m, v815, v819, v811)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L18
	} else {
		goto L258
	}
L258:
	;
	goto L222
L259:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v595)+56))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v834+v621)+4))
	if v836 != 0 {
		v628 = v836
		goto L206
	} else {
		goto L260
	}
L260:
	;
	goto L207
L261:
	;
	goto L203
L262:
	;
	goto L198
L263:
	;
	v880 = int32(_a_F_ProcessInterrupts_14)
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v882 - int32(1)
	m.G0 = v553 + int32(160)
	goto L187
L264:
	;
	F_ProcessLogMemoryContextInterrupt(m)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L18
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v904 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[47]))
	if v904 == int32(0) {
		goto L12
	} else {
		goto L268
	}
L267:
	;
	goto L266
L268:
	;
	v907 = m.G0
	v909 = v907 - int32(176)
	m.G0 = v909
	v911 = int32(_a_F_ProcessInterrupts_14)
	v913 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v913 + int32(1)
	v918 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[48]))
	if v918 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v936 = int32(0)
	v937 = int32(_a_F_ProcessInterrupts_15)
	v938 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40])) = v935
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[47])) = v936
	v945 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[49]))
	if v945 != 0 {
		goto L275
	} else {
		goto L276
	}
L270:
	;
	v923 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[42]))
	v928 = F_AllocSetContextCreateInternal(m, v923, int32(_a_F_ProcessInterrupts_28), int32(0), int32(_a_F_ProcessInterrupts_17), int32(_a_F_ProcessInterrupts_18))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L18
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	F_MemoryContextReset(m, v918)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L18
	} else {
		goto L274
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[48])) = v928
	v935 = v928
	goto L269
L274:
	;
	v934 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[48]))
	v935 = v934
	goto L269
L275:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v945)+4))
	if int32(0) < v946 {
		goto L278
	} else {
		goto L279
	}
L276:
	;
	v1093 = v935
	goto L277
L277:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40])) = v938
	F_MemoryContextReset(m, v1093)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L18
	} else {
		goto L316
	}
L278:
	;
	v951 = v936
	goto L281
L279:
	;
	goto L280
L280:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[48]))
	v1093 = v1091
	goto L277
L281:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v945)+12))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v959+v951<<(uint(int32(2))%32))))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v963)+4))
	if v964 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	goto L280
L283:
	;
	v1077 = v951 + int32(1)
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v945)+4))
	if v1077 < v1078 {
		v951 = v1077
		goto L281
	} else {
		goto L315
	}
L284:
	;
	v972 = F_shm_mq_receive(m, v964, v909+int32(72), v909+int32(68), int32(1))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L18
	} else {
		goto L288
	}
L285:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v909)+52))
	F_pfree(m, v1072)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L18
	} else {
		goto L314
	}
L286:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L18
	} else {
		goto L310
	}
L287:
	;
	F_initStringInfo(m, v909+int32(52))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L18
	} else {
		goto L289
	}
L288:
	;
	switch v972 {
	case 0:
		goto L287
	case 1:
		goto L283
	default:
		goto L286
	}
L289:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v909)+68))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v909)+72))
	F_appendBinaryStringInfo(m, v909+int32(52), v980, v981)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L18
	} else {
		goto L290
	}
L290:
	;
	v986 = F_pq_getmsgbyte(m, v909+int32(52))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L18
	} else {
		goto L293
	}
L291:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L18
	} else {
		goto L307
	}
L292:
	;
	F_pq_parse_errornotice(m, v909+int32(52), v909+int32(76))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L18
	} else {
		goto L294
	}
L293:
	;
	v988 = base.I32_extend8_s(v986)
	switch v988 - int32(65) {
	case 0, 13:
		goto L285
	default:
		goto L291
	case 4:
		goto L292
	}
L294:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v909)+124))
	if v997 != 0 {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[50]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[44])) = v1012
	*(*int32)(unsafe.Add(mBase, uint32(v909)+124)) = v1009
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L18
	} else {
		goto L301
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v909)+36)) = int32(_a_F_ProcessInterrupts_29)
	*(*int32)(unsafe.Add(mBase, uint32(v909)+32)) = v997
	v1004 = F_psprintf(m, int32(_a_F_ProcessInterrupts_24), v909+int32(32))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L18
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v1007 = F_pstrdup(m, int32(_a_F_ProcessInterrupts_29))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L18
	} else {
		goto L300
	}
L299:
	;
	v1009 = v1004
	goto L295
L300:
	;
	v1009 = v1007
	goto L295
L301:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L18
	} else {
		goto L302
	}
L302:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_30), int32(0))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L18
	} else {
		goto L303
	}
L303:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L18
	} else {
		goto L304
	}
L304:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v909)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v909)+16)) = v1029
	F_errcontext_msg(m, int32(_a_F_ProcessInterrupts_31), v909+int32(16))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L18
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_32), int32(1048), int32(_a_F_ProcessInterrupts_33))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L18
	} else {
		goto L306
	}
L306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v909))) = v988
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v909)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v909)+4)) = v1046
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_34), v909)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L18
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_32), int32(1062), int32(_a_F_ProcessInterrupts_33))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L18
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L18
	} else {
		goto L311
	}
L311:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_35), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L18
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_32), int32(1134), int32(_a_F_ProcessInterrupts_28))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L18
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	goto L283
L315:
	;
	goto L282
L316:
	;
	v1106 = int32(_a_F_ProcessInterrupts_14)
	v1108 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v1108 - int32(1)
	m.G0 = v909 + int32(176)
	goto L12
L317:
	;
	F_errcode(m, int32(50463042))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L18
	} else {
		goto L318
	}
L318:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_36), int32(0))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L18
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3593), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L18
	} else {
		goto L320
	}
L320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L321:
	;
	F_errcode(m, int32(67240258))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L18
	} else {
		goto L322
	}
L322:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_37), int32(0))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L18
	} else {
		goto L323
	}
L323:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3606), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L18
	} else {
		goto L324
	}
L324:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L325:
	;
	F_errcode(m, int32(84017605))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L18
	} else {
		goto L326
	}
L326:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_38), int32(0))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L18
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3619), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L18
	} else {
		goto L328
	}
L328:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L329:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L18
	} else {
		goto L330
	}
L330:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L18
	} else {
		goto L331
	}
L331:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_39), int32(0))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L18
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3572), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L18
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L18
	} else {
		goto L335
	}
L335:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L18
	} else {
		goto L336
	}
L336:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_40), int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L18
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3559), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L18
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L18
	} else {
		goto L340
	}
L340:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L18
	} else {
		goto L341
	}
L341:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_41), int32(0))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L18
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3552), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L18
	} else {
		goto L343
	}
L343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L344:
	;
	v1236 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[7])) = v1236
	F_errstart_cold(m, int32(22), v1236)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L18
	} else {
		goto L345
	}
L345:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L18
	} else {
		goto L346
	}
L346:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_42), int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L18
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3493), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L18
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L18
	} else {
		goto L350
	}
L350:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_43), int32(0))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L18
	} else {
		goto L351
	}
L351:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3431), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L18
	} else {
		goto L352
	}
L352:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L353:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L18
	} else {
		goto L354
	}
L354:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_44), int32(0))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L18
	} else {
		goto L355
	}
L355:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3427), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L18
	} else {
		goto L356
	}
L356:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L357:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L18
	} else {
		goto L358
	}
L358:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_45), int32(0))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L18
	} else {
		goto L359
	}
L359:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3423), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L18
	} else {
		goto L360
	}
L360:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L361:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L18
	} else {
		goto L362
	}
L362:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L18
	} else {
		goto L363
	}
L363:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_46), int32(0))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L18
	} else {
		goto L364
	}
L364:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3545), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L18
	} else {
		goto L365
	}
L365:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessNotifyInterrupt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	goto L2
L1:
	;
	return
L2:
	;
	if v4 != int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[1]))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L5
L5:
	;
	v13 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[1])) = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[2]))
	if v16 == v13 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L1
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[1]))
	if v69 != 0 {
		goto L5
	} else {
		goto L28
	}
L8:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[3])))
	if v20 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L16
	}
L10:
	;
	v25 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	if v25 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessNotifyInterrupt_0), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_ProcessNotifyInterrupt_1), int32(2327), int32(_a_F_ProcessNotifyInterrupt_0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	F_asyncQueueReadAllNotifications(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	if l0 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[4]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = m.T0[v46].(func(*base.Module) int32)(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[3])))
	if v50 != int32(1) {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v55 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	if v55 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessNotifyInterrupt_2), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_ProcessNotifyInterrupt_1), int32(2351), int32(_a_F_ProcessNotifyInterrupt_0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	goto L7
L28:
	;
	goto L6
}
func F_PromoteIsTriggered(m *base.Module) int32 {
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[0])))
	if v5 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = int32(1)
		if v10 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
			F_s_lock(m, v14+int32(96), int32(_a_F_PromoteIsTriggered_0), int32(_a_F_PromoteIsTriggered_1), int32(_a_F_PromoteIsTriggered_2))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
				*(*uint8)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[0])) = uint8(v27)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = int32(0)
				v31 = v27
				return v31 & int32(1)
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
			*(*uint8)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[0])) = uint8(v27)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = int32(0)
			v31 = v27
			return v31 & int32(1)
		}
	} else {
		v31 = int32(1)
		return v31 & int32(1)
	}
}
func F_p_isasclet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v6 != int32(1) {
		v78 = int32(0)
		return v78
	} else {
		v9 = int32(0)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v13 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10+v11))))
		if v13 < v9 {
			v78 = v9
			return v78
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
			if v16 == int32(1) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v19 != 0 {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20<<(uint(int32(2))%32))))
					return base.B2i32(base.Ui32(int32(127)) < base.Ui32(v24)) | base.B2i32(base.Ui32(v24|int32(32)-int32(97)) < base.Ui32(int32(26)))
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36<<(uint(int32(2))%32))))
					if base.Ui32(v40) <= base.Ui32(int32(_a_F_p_isasclet_0)) {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v40)>>(uint(int32(8))%32)))+uint32(_c_F_p_isasclet[0]))))
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v40)>>(uint(int32(3))%32))&int32(31)|v51<<(uint(int32(5))%32))+uint32(_c_F_p_isasclet[0]))))
						v65 = int32(base.Ui32(v57)>>(uint(v40&int32(7))%32)) & int32(1)
					} else {
						v65 = base.B2i32(base.Ui32(v40) < base.Ui32(int32(_a_F_p_isasclet_1)))
					}
					return base.B2i32(v65 != int32(0))
				}
			} else {
				v78 = base.B2i32(base.Ui32((v13|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)))
				return v78
			}
		}
	}
}
func F_p_isnotalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v68 int32
	_ = v68
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6+v9<<(uint(int32(2))%32))))
			if base.Ui32(int32(127)) < base.Ui32(v13) {
				v68 = int32(1)
				return base.B2i32(v68 == int32(0))
			} else {
				return base.B2i32(base.B2i32(base.Ui32(v13-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v13|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+v32<<(uint(int32(2))%32))))
			if base.Ui32(int32(10)) <= base.Ui32(v36-int32(48)) {
				v43 = F_iswalpha(m, v36)
				mBase = m.M
				v46 = base.B2i32(v43 != int32(0))
			} else {
				v46 = int32(1)
			}
			return base.B2i32(v46 == int32(0))
		}
	} else {
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
		v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v52))))
		v68 = base.B2i32(base.Ui32(v54-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v54|int32(32)-int32(97)) < base.Ui32(int32(26)))
		return base.B2i32(v68 == int32(0))
	}
}
func F_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v7 = m.G0
	v9 = v7 - int32(256)
	m.G0 = v9
	if l2 <= l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(256)
	return
L2:
	;
	if l4&int32(_a_F_pad_0) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = l2 - l3
	v15 = int32(256)
	v17 = base.B2i32(base.Ui32(v14) < base.Ui32(v15))
	if base.Ui32(v14) < base.Ui32(v15) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v14
	goto L6
L5:
	;
	v18 = v15
	goto L6
L6:
	;
	v20 = F__emscripten_memset_bulkmem(m, v9, base.I32_extend8_s(l1), v18)
	goto L7
L7:
	;
	if v17 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = v14
	goto L11
L9:
	;
	v39 = v14
	goto L10
L10:
	;
	F_out(m, l0, v9, v39)
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L16
	}
L11:
	;
	F_out(m, l0, v9, int32(256))
	v31 = m.ExcPending
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v39 = v33
	goto L10
L13:
	;
	return
L14:
	;
	v33 = v26 - int32(256)
	if base.Ui32(int32(255)) < base.Ui32(v33) {
		v26 = v33
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	goto L1
}
func F_pair_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v37 int32
	_ = v37
	var v41 float64
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = l0
	goto L1
L1:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v23-int32(9)))&base.B2i32(v23 != int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v37 = v14 + base.B2i32(v23 == int32(40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v37
	v41 = F_float8in_internal(m, v37, v12+int32(12), l4, l5, l6)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v14 = v14 + int32(1)
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	return int32(0)
L7:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v41
	if l6 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v12 + int32(16)
	return v146
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v57 = v55 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != int32(44) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v48 != int32(447) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v51 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v146 = int32(0)
	goto L8
L13:
	;
	v128 = int32(0)
	v129 = F_errsave_start(m, l6)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L30
	}
L14:
	;
	v64 = F_float8in_internal(m, v57, v12+int32(12), l4, l5, l6)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = v64
	if l6 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v23 != int32(40) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v69 != int32(447) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v72 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v146 = int32(0)
	goto L8
L20:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if l3 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v80 = v78 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 != int32(41) {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v85 = v80
	goto L23
L23:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v94-int32(9)))&base.B2i32(v94 != int32(32)) != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v103 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v103
	v85 = v103
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v114
	v146 = int32(1)
	goto L8
L27:
	;
	goto L28
L28:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v117 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v146 = int32(1)
	goto L8
L30:
	;
	if v129 == int32(0) {
		v146 = v128
		goto L8
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l4
	F_errmsg(m, int32(_a_F_pair_decode_0), v12)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	F_errsave_finish(m, l6, int32(_a_F_pair_decode_1), int32(251), int32(_a_F_pair_decode_2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v146 = v128
	goto L8
}
func F_pairingheap_remove(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = F_pairingheap_remove_first(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v20 = v15 + base.B2i32(v16 != l1)<<(uint(int32(2))%32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	return
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v21
	if v21 == int32(0) {
		goto L6
	} else {
		goto L46
	}
L10:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v91
	if v21 == int32(0) {
		goto L6
	} else {
		goto L45
	}
L11:
	;
	v91 = v22
	goto L10
L12:
	;
	goto L13
L13:
	;
	v29 = int32(0)
	v30 = v22
	goto L14
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v37 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v29 == int32(0) {
		v91 = v57
		goto L10
	} else {
		goto L31
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v29
	v57 = v30
	goto L16
L18:
	;
	goto L19
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = m.T0[v43].(func(*base.Module, int32, int32, int32) int32)(m, v30, v37, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v47 = base.B2i32(v44 < int32(0))
	if v44 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v48 = v30
	goto L23
L22:
	;
	v48 = v37
	goto L23
L23:
	;
	if v44 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v49 = v37
	goto L26
L25:
	;
	v49 = v30
	goto L26
L26:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v48
	goto L29
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v49
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v48
	if v41 != 0 {
		v29 = v49
		v30 = v41
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v57 = v49
	goto L16
L31:
	;
	v66 = v57
	v68 = v29
	goto L32
L32:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v77 = m.T0[v76].(func(*base.Module, int32, int32, int32) int32)(m, v66, v68, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	v91 = v82
	goto L10
L34:
	;
	v80 = base.B2i32(v77 < int32(0))
	if v77 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v81 = v66
	goto L37
L36:
	;
	v81 = v68
	goto L37
L37:
	;
	if v77 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v82 = v68
	goto L40
L39:
	;
	v82 = v66
	goto L40
L40:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v83 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v81
	goto L43
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v82
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v81
	if v74 != 0 {
		v66 = v82
		v68 = v74
		goto L32
	} else {
		goto L44
	}
L44:
	;
	goto L33
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v91
	return
L46:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v110
	goto L6
}
func F_paramlist_parser_setup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(815)
	return
}
func F_parse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v372 int32
	_ = v372
	v6 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = m.T0[v16].(func(*base.Module) int32)(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if v13 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v25 = v23
	goto L8
L7:
	;
	v25 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v25
	return int32(0)
L9:
	;
	v49 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v48)+4)) = v49
	v51 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)) = uint8(v51)
	v53 = int32(380)
	*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v48)+20)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(281479271677952)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v63 != 0 {
		v372 = v6
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v29
	v48 = v13
	goto L9
L11:
	;
	goto L12
L12:
	;
	v33 = F_palloc_extended(m, int32(88), int32(2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v33 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v39 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v33
	v48 = v33
	goto L9
L17:
	;
	v41 = v39
	goto L19
L18:
	;
	v41 = int32(12)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v41
	return int32(0)
L20:
	;
	return v372
L21:
	;
	v65 = v48 + int32(20)
	v77 = v6
	goto L23
L22:
	;
	if l1 != v278 {
		goto L94
	} else {
		goto L95
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v79 = F_newstate(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v278 = v277
	goto L22
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v82 = F_newstate(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_parse[0]))
	if v89 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v92 <= v93 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L32
L34:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_parse[0]))
	if v169 != 0 {
		goto L56
	} else {
		goto L57
	}
L35:
	;
	F_createarc(m, v87, int32(110), int32(0), l3, v79)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L55
	}
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v95 == int32(0) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	if v117 == int32(0) {
		goto L35
	} else {
		goto L47
	}
L39:
	;
	v103 = v95
	goto L40
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	if v110 != v79 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L35
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	if v116 != 0 {
		v103 = v116
		goto L40
	} else {
		goto L46
	}
L43:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
	if v112 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v113 == int32(110) {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	goto L41
L47:
	;
	v125 = v117
	goto L48
L48:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	if v132 != l3 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L35
L50:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v125)+24))
	if v138 != 0 {
		v125 = v138
		goto L48
	} else {
		goto L54
	}
L51:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+4)))
	if v134 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v135 == int32(110) {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	goto L49
L55:
	;
	goto L34
L56:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v172 <= v173 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L58
L60:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v247 != 0 {
		goto L82
	} else {
		goto L83
	}
L61:
	;
	F_createarc(m, v167, int32(110), int32(0), v82, l4)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L81
	}
L62:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v175 == int32(0) {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v197 == int32(0) {
		goto L61
	} else {
		goto L73
	}
L65:
	;
	v183 = v175
	goto L66
L66:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	if v190 != l4 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L61
L68:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	if v196 != 0 {
		v183 = v196
		goto L66
	} else {
		goto L72
	}
L69:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+4)))
	if v192 != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v193 == int32(110) {
		goto L60
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	goto L67
L73:
	;
	v205 = v197
	goto L74
L74:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	if v212 != v82 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L61
L76:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v205)+24))
	if v218 != 0 {
		v205 = v218
		goto L74
	} else {
		goto L80
	}
L77:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+4)))
	if v214 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if v215 == int32(110) {
		goto L60
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	goto L75
L81:
	;
	goto L60
L82:
	;
	return int32(0)
L83:
	;
	goto L84
L84:
	;
	v250 = int32(0)
	v252 = F_parsebranch(m, l0, l1, l2, v79, v82, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v254 != 0 {
		v372 = v250
		goto L20
	} else {
		goto L86
	}
L86:
	;
	if v77 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+1)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v261 = v257 | v260
	v270 = v257&int32(28) | v261<<(uint(int32(1))%32)&(v261<<(uint(int32(2))%32))&int32(4) | v260
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v270)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v272 != int32(124) {
		v278 = v272
		goto L22
	} else {
		goto L91
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = v252
	goto L87
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v252
	goto L87
L91:
	;
	v275 = F_next(m, l0)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v275 != 0 {
		v77 = v252
		goto L23
	} else {
		goto L93
	}
L93:
	;
	goto L24
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v283 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v288 == v252 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v285 = v283
	goto L99
L98:
	;
	v285 = int32(8)
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v285
	goto L96
L100:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
	if v290 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v316&int32(28) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L103:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v48)+64))
	F_pfree(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v302 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v302)
	v304 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v304
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v304
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v308 != 0 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	F_pfree(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
	F_pfree(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = int32(0)
	goto L105
L109:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v48
	return v252
L110:
	;
	goto L111
L111:
	;
	F_pfree(m, v48)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	return v252
L113:
	;
	if v288 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	v372 = v48
	goto L20
L116:
	;
	v328 = v288
	goto L119
L117:
	;
	goto L118
L118:
	;
	v348 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v348)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(0)
	goto L115
L119:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	F_freesubre(m, l0, v328)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	goto L118
L121:
	;
	if v333 != 0 {
		v328 = v333
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
}
func F_parseCheckAggregates(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v675 int32
	_ = v675
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int64
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int64
	_ = v736
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v3
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v25 == v3 {
		v253 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v261 == int32(0) {
		v372 = v3
		v373 = v3
		goto L42
	} else {
		goto L43
	}
L2:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+104)))
	v30 = F_expand_grouping_sets(m, v25, v28, int32(_a_F_parseCheckAggregates_0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v30 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v57 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	F_errcode(m, int32(16777477))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(_a_F_parseCheckAggregates_1), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = v45
	goto L13
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v47 = v46
	goto L13
L13:
	;
	v48 = F_exprLocation(m, v47)
	mBase = m.M
	F_parser_errposition(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_parseCheckAggregates_2), int32(1172), int32(_a_F_parseCheckAggregates_3))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v235 != int32(1) {
		v253 = v227
		goto L1
	} else {
		goto L40
	}
L17:
	;
	v227 = int32(0)
	goto L16
L18:
	;
	v60 = int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v61 <= v60 {
		v227 = v57
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v68 = v60
	v74 = v57
	goto L20
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v68<<(uint(int32(2))%32))))
	v87 = int32(0)
	if v74 == v87 {
		v182 = v87
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v227 = v182
	goto L16
L22:
	;
	if v182 == int32(0) {
		goto L17
	} else {
		goto L38
	}
L23:
	;
	if v86 == int32(0) {
		v182 = v87
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v93 <= int32(0) {
		v182 = v87
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v104 = v87
	v105 = v87
	v107 = v93
	goto L26
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v114 <= int32(0) {
		v161 = v104
		v164 = v107
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v182 = v161
	goto L22
L28:
	;
	v172 = v105 + int32(1)
	if v172 < v164 {
		v104 = v161
		v105 = v172
		v107 = v164
		goto L26
	} else {
		goto L37
	}
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v105<<(uint(int32(2))%32))))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v129 = int32(0)
	goto L30
L30:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v122+v129<<(uint(int32(2))%32))))
	if v121 != v145 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v150 = F_lappend_int(m, v104, v121)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L36
	}
L32:
	;
	v148 = v129 + int32(1)
	if v148 != v114 {
		v129 = v148
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	v161 = v104
	v164 = v107
	goto L28
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v161 = v150
	v164 = v152
	goto L28
L37:
	;
	goto L27
L38:
	;
	v195 = v68 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v195 < v196 {
		v68 = v195
		v74 = v182
		goto L20
	} else {
		goto L39
	}
L39:
	;
	goto L21
L40:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v238 == int32(0) {
		v253 = v227
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = int32(0)
	v253 = v227
	goto L1
L42:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v377 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L43:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v264 <= int32(0) {
		v372 = v3
		v373 = v3
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v267 = int32(0)
	if v267 < v264 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v270 = v264
	goto L47
L46:
	;
	v270 = v267
	goto L47
L47:
	;
	v271 = int32(1)
	if v264 == v271 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v270&v271 == int32(0) {
		v372 = v330
		v373 = v341
		goto L42
	} else {
		goto L61
	}
L49:
	;
	v275 = int32(0)
	v330 = v275
	v331 = v275
	v341 = v3
	goto L48
L50:
	;
	goto L51
L51:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v285 = int32(0)
	v288 = v3
	v295 = v3
	v298 = v3
	goto L52
L52:
	;
	v299 = int32(1)
	v301 = int32(2)
	v303 = v279 + v285<<(uint(v301)%32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	switch v305 - v301 {
	case 0:
		v311 = v299
		v312 = v295
		goto L54
	default:
		v310 = v295
		goto L55
	case 4:
		goto L56
	}
L53:
	;
	v330 = v320
	v331 = v323
	v341 = v321
	goto L48
L54:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+12))
	switch v314 - int32(2) {
	case 0:
		v320 = v299
		v321 = v312
		goto L57
	default:
		v319 = v312
		goto L58
	case 4:
		goto L59
	}
L55:
	;
	v311 = v298
	v312 = v310
	goto L54
L56:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+92)))
	v310 = v308 | v295
	goto L55
L57:
	;
	v322 = int32(2)
	v323 = v285 + v322
	v325 = v288 + v322
	if v325 != v270&int32(2147483646) {
		v285 = v323
		v288 = v325
		v295 = v321
		v298 = v320
		goto L52
	} else {
		goto L60
	}
L58:
	;
	v320 = v311
	v321 = v319
	goto L57
L59:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+92)))
	v319 = v317 | v312
	goto L58
L60:
	;
	goto L53
L61:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v349 = int32(2)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348+v331<<(uint(v349)%32))))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+12))
	switch v353 - v349 {
	case 0:
		v372 = int32(1)
		v373 = v341
		goto L42
	default:
		v358 = v341
		goto L62
	case 4:
		goto L63
	}
L62:
	;
	v372 = v330
	v373 = v358
	goto L42
L63:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+92)))
	v358 = v356 | v341
	goto L62
L64:
	;
	if v372&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L65:
	;
	v424 = int32(0)
	goto L64
L66:
	;
	v380 = int32(0)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	if v381 <= v380 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v388 = v380
	v389 = int32(0)
	goto L68
L68:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v377)+12))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403+v388<<(uint(int32(2))%32))))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v409 = F_get_sortgroupclause_tle(m, v407, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L3
	} else {
		goto L70
	}
L69:
	;
	v424 = v413
	goto L64
L70:
	;
	if v409 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v411 = F_lappend(m, v389, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L3
	} else {
		goto L74
	}
L72:
	;
	v413 = v389
	goto L73
L73:
	;
	v415 = v388 + int32(1)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	if v415 < v416 {
		v388 = v415
		v389 = v413
		goto L68
	} else {
		goto L75
	}
L74:
	;
	v413 = v411
	goto L73
L75:
	;
	goto L69
L76:
	;
	v441 = F_flatten_join_alias_vars(m, int32(0), l1, v424)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L3
	} else {
		goto L79
	}
L77:
	;
	v443 = v424
	goto L78
L78:
	;
	if v443 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v443 = v441
	goto L78
L80:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v689 = v21 + int32(44)
	v690 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v689))) = uint8(v690)
	v692 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+36)) = v692
	v694 = int32(1)
	v695 = v686 & v694
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v695)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v692
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v443
	v701 = v372 & v694
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)) = uint8(v701)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	v707 = F_finalize_grouping_exprs_walker(m, v687, v21+int32(8))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L3
	} else {
		goto L140
	}
L81:
	;
	v446 = int32(0)
	v675 = v446
	v686 = v446
	goto L80
L82:
	;
	goto L83
L83:
	;
	v448 = int32(0)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v449 <= v448 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v554 = int32(0)
	v560 = F_palloc0(m, int32(136))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L3
	} else {
		goto L113
	}
L85:
	;
	v542 = int32(0)
	v553 = v448
	goto L84
L86:
	;
	goto L87
L87:
	;
	v453 = int32(0)
	v458 = v453
	v461 = v453
	v472 = v448
	goto L88
L88:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v443)+12))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v473+v458<<(uint(int32(2))%32))))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	if v479 != int32(6) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v542 = v530
	v553 = v531
	goto L84
L90:
	;
	v533 = v458 + int32(1)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v533 < v534 {
		v458 = v533
		v461 = v530
		v472 = v531
		goto L88
	} else {
		goto L112
	}
L91:
	;
	v530 = v461
	v531 = int32(1)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v483 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v477)+16))
	v485 = int32(0)
	if v253 == v485 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v527 = v478
	goto L96
L96:
	;
	v528 = F_lappend(m, v461, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L3
	} else {
		goto L111
	}
L97:
	;
	if v523 == int32(0) {
		v530 = v461
		v531 = v472
		goto L90
	} else {
		goto L110
	}
L98:
	;
	v523 = int32(0)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v491 <= int32(0) {
		v516 = v485
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v523 = v516
	goto L97
L102:
	;
	v494 = int32(0)
	if v494 < v491 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v497 = v491
	goto L105
L104:
	;
	v497 = v494
	goto L105
L105:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v500 = int32(0)
	goto L106
L106:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v498+v500<<(uint(int32(2))%32))))
	v509 = base.B2i32(v508 == v484)
	if v508 == v484 {
		v516 = v509
		goto L101
	} else {
		goto L108
	}
L107:
	;
	v516 = v509
	goto L101
L108:
	;
	v511 = v500 + int32(1)
	if v511 != v497 {
		v500 = v511
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	v527 = v526
	goto L96
L111:
	;
	v530 = v528
	v531 = v472
	goto L90
L112:
	;
	goto L89
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v560)+12)) = int32(9)
	*(*int64)(unsafe.Add(mBase, uint32(v560))) = int64(101)
	v568 = F_makeAlias(m, int32(_a_F_parseCheckAggregates_4), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	if v443 == int32(0) {
		v642 = v554
		v644 = v554
		v645 = v554
		v648 = v554
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v651 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v560)+124)) = uint16(v651)
	*(*int32)(unsafe.Add(mBase, uint32(v560)+120)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v560)+8)) = v568
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v656 = F_lappend(m, v655, v560)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L3
	} else {
		goto L135
	}
L116:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v572 <= int32(0) {
		v642 = v554
		v644 = v554
		v645 = v554
		v648 = v554
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v584 = v554
	v586 = v554
	v587 = v554
	v590 = v554
	v591 = v554
	goto L118
L118:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v443)+12))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v593+v591<<(uint(int32(2))%32))))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	if v598 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v642 = v622
	v644 = v617
	v645 = v612
	v648 = v627
	goto L115
L120:
	;
	v599 = F_pstrdup(m, v598)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L3
	} else {
		goto L123
	}
L121:
	;
	v602 = int32(_a_F_parseCheckAggregates_5)
	goto L122
L122:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v568)+8))
	v604 = F_makeString(m, v602)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L3
	} else {
		goto L124
	}
L123:
	;
	v602 = v599
	goto L122
L124:
	;
	v606 = F_lappend(m, v603, v604)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v568)+8)) = v606
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	v610 = F_copyObjectImpl(m, v609)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	v612 = F_lappend(m, v587, v610)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	v615 = F_exprType(m, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	v617 = F_lappend_oid(m, v586, v615)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	v620 = F_exprTypmod(m, v619)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	v622 = F_lappend_int(m, v584, v620)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	v625 = F_exprCollation(m, v624)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	v627 = F_lappend_oid(m, v590, v625)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	v630 = v591 + int32(1)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v630 < v631 {
		v584 = v622
		v586 = v617
		v587 = v612
		v590 = v627
		v591 = v630
		goto L118
	} else {
		goto L134
	}
L134:
	;
	goto L119
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v656
	if v656 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	v661 = v659
	goto L138
L137:
	;
	v661 = int32(0)
	goto L138
L138:
	;
	v662 = F_buildNSItemFromLists(m, v560, v661, v644, v642, v648)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L3
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v662
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v666 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v666)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v665
	v675 = v542
	v686 = v553
	goto L80
L140:
	;
	if v701 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v710 = F_flatten_join_alias_vars(m, int32(0), l1, v687)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L3
	} else {
		goto L144
	}
L142:
	;
	v712 = v687
	goto L143
L143:
	;
	v713 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v713)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v713
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v695)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v443
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)) = uint8(v713)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(4)
	v730 = F_substitute_grouped_columns_mutator(m, v712, v21+int32(8))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L3
	} else {
		goto L145
	}
L144:
	;
	v712 = v710
	goto L143
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v730
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v734 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v689))) = uint8(v734)
	v736 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+36)) = v736
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v695)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v443
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)) = uint8(v701)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	v747 = F_finalize_grouping_exprs_walker(m, v733, v21+int32(8))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L3
	} else {
		goto L146
	}
L146:
	;
	if v701 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v750 = F_flatten_join_alias_vars(m, int32(0), l1, v733)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L3
	} else {
		goto L150
	}
L148:
	;
	v752 = v733
	goto L149
L149:
	;
	v753 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v753)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v753
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v695)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v443
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)) = uint8(v753)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	v770 = F_substitute_grouped_columns_mutator(m, v752, v21+int32(8))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L3
	} else {
		goto L151
	}
L150:
	;
	v752 = v750
	goto L149
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v770
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v773&v373&int32(1) != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L3
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	m.G0 = v21 + int32(48)
	return
L155:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L3
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_parseCheckAggregates_6), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	v789 = F_locate_agg_of_level(m, l1, int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L3
	} else {
		goto L158
	}
L158:
	;
	F_parser_errposition(m, l0, v789)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L3
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_parseCheckAggregates_2), int32(1331), int32(_a_F_parseCheckAggregates_3))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L3
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_datetime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v117 int32
	_ = v117
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int64
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v426 int64
	_ = v426
	var v429 int64
	_ = v429
	var v430 int64
	_ = v430
	var v433 int64
	_ = v433
	var v434 int64
	_ = v434
	var v436 int64
	_ = v436
	var v437 int64
	_ = v437
	var v440 int64
	_ = v440
	var v449 int32
	_ = v449
	var v450 int64
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v479 int64
	_ = v479
	var v482 int64
	_ = v482
	var v483 int64
	_ = v483
	var v486 int64
	_ = v486
	var v487 int64
	_ = v487
	var v489 int64
	_ = v489
	var v490 int64
	_ = v490
	var v493 int64
	_ = v493
	var v501 int64
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v27 = F_do_to_timestamp(m, l0, l1, int32(100), int32(1), v13+int32(52), v13+int32(40), v13+int32(44), v13+int32(36), v13+int32(32), l5)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		if v27 == int32(0) {
			v581 = int32(0)
			m.G0 = v13 + int32(96)
			return v581
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
			if v33 != 0 {
				v35 = v33
			} else {
				v35 = int32(-1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
			v39 = v37 & int32(2)
			if v37&int32(1) != 0 {
				v43 = v37 & int32(4)
				if v39 != 0 {
					if v43 != 0 {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+44)))
						if v44 == int32(1) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v47
							v50 = v13 + int32(52)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
							v53 = v13 + int32(24)
							v60 = m.G0
							v62 = v60 - int32(16)
							m.G0 = v62
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
							if v64 <= int32(-4713) {
								if v64 != int32(-4713) {
									*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
									v141 = int32(-1)
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
									if int32(10) < v69 {
										v80 = v69
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
										v82 = F_date2j(m, v64, v80, v81)
										mBase = m.M
										v85 = base.I64_extend_i32_s(v82 - int32(_a_F_parse_datetime_0))
										v86 = int64(63)
										F___multi3(m, v62, v85, v85>>(uint(v86)%64), int64(86400000000), int64(0))
										mBase = m.M
										v91 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
										v92 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
										if v91 != v92>>(uint(v86)%64) {
											*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
											v141 = int32(-1)
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
											v100 = int32(60)
											v109 = base.I64_extend_i32_s(v51) + base.I64_extend_i32_s(v97+(v98+v99*v100)*v100)*int64(1000000)
											v110 = v92 + v109
											*(*int64)(unsafe.Add(mBase, uint32(v53))) = v110
											if base.B2i32(v109 < int64(0))^base.B2i32(v110 < v92) != 0 {
												*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
												v141 = int32(-1)
											} else {
												if l4 != 0 {
													v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
													v122 = base.I64_extend_i32_s(int32(0)-v117)*int64(-1000000) + v110
													*(*int64)(unsafe.Add(mBase, uint32(v53))) = v122
													v124 = v122
												} else {
													v124 = v110
												}
												if base.Ui64(v124+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
													v141 = int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
													v141 = int32(-1)
												}
											}
										}
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
										v141 = int32(-1)
									}
								}
							} else {
								if v64 <= int32(_a_F_parse_datetime_1) {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
									v80 = v74
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
									v82 = F_date2j(m, v64, v80, v81)
									mBase = m.M
									v85 = base.I64_extend_i32_s(v82 - int32(_a_F_parse_datetime_0))
									v86 = int64(63)
									F___multi3(m, v62, v85, v85>>(uint(v86)%64), int64(86400000000), int64(0))
									mBase = m.M
									v91 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
									v92 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
									if v91 != v92>>(uint(v86)%64) {
										*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
										v141 = int32(-1)
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
										v100 = int32(60)
										v109 = base.I64_extend_i32_s(v51) + base.I64_extend_i32_s(v97+(v98+v99*v100)*v100)*int64(1000000)
										v110 = v92 + v109
										*(*int64)(unsafe.Add(mBase, uint32(v53))) = v110
										if base.B2i32(v109 < int64(0))^base.B2i32(v110 < v92) != 0 {
											*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
											v141 = int32(-1)
										} else {
											if l4 != 0 {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
												v122 = base.I64_extend_i32_s(int32(0)-v117)*int64(-1000000) + v110
												*(*int64)(unsafe.Add(mBase, uint32(v53))) = v122
												v124 = v122
											} else {
												v124 = v110
											}
											if base.Ui64(v124+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v141 = int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
												v141 = int32(-1)
											}
										}
									}
								} else {
									if v64 != int32(_a_F_parse_datetime_2) {
										*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
										v141 = int32(-1)
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
										if int32(5) < v77 {
											*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
											v141 = int32(-1)
										} else {
											v80 = v77
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
											v82 = F_date2j(m, v64, v80, v81)
											mBase = m.M
											v85 = base.I64_extend_i32_s(v82 - int32(_a_F_parse_datetime_0))
											v86 = int64(63)
											F___multi3(m, v62, v85, v85>>(uint(v86)%64), int64(86400000000), int64(0))
											mBase = m.M
											v91 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
											v92 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
											if v91 != v92>>(uint(v86)%64) {
												*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
												v141 = int32(-1)
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
												v100 = int32(60)
												v109 = base.I64_extend_i32_s(v51) + base.I64_extend_i32_s(v97+(v98+v99*v100)*v100)*int64(1000000)
												v110 = v92 + v109
												*(*int64)(unsafe.Add(mBase, uint32(v53))) = v110
												if base.B2i32(v109 < int64(0))^base.B2i32(v110 < v92) != 0 {
													*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
													v141 = int32(-1)
												} else {
													if l4 != 0 {
														v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
														v122 = base.I64_extend_i32_s(int32(0)-v117)*int64(-1000000) + v110
														*(*int64)(unsafe.Add(mBase, uint32(v53))) = v122
														v124 = v122
													} else {
														v124 = v110
													}
													if base.Ui64(v124+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
														v141 = int32(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v53))) = int64(0)
														v141 = int32(-1)
													}
												}
											}
										}
									}
								}
							}
							m.G0 = v62 + int32(16)
							if v141 == int32(0) {
								v183 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								F_AdjustTimestampForTypmod(m, v13+int32(24), v183, l5)
								mBase = m.M
								v185 = m.ExcPending
								if v185 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1184)
									v188 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
									v189 = F_Int64GetDatum(m, v188)
									mBase = m.M
									v190 = m.ExcPending
									if v190 != 0 {
										return int32(0)
									} else {
										v581 = v189
										m.G0 = v13 + int32(96)
										return v581
									}
								}
							} else {
								v147 = int32(0)
								v148 = F_errsave_start(m, l5)
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									if v148 == int32(0) {
										v581 = v147
										m.G0 = v13 + int32(96)
										return v581
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_parse_datetime_3), int32(0))
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_5), int32(_a_F_parse_datetime_6))
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int32(0)
												} else {
													v581 = v147
													m.G0 = v13 + int32(96)
													return v581
												}
											}
										}
									}
								}
							}
						} else {
							v164 = int32(0)
							v165 = F_errsave_start(m, l5)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								if v165 == int32(0) {
									v581 = v164
									m.G0 = v13 + int32(96)
									return v581
								} else {
									F_errcode(m, int32(117440642))
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_parse_datetime_7), int32(0))
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_8), int32(_a_F_parse_datetime_6))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return int32(0)
											} else {
												v581 = v164
												m.G0 = v13 + int32(96)
												return v581
											}
										}
									}
								}
							}
						}
					} else {
						v191 = int32(0)
						v193 = v13 + int32(52)
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
						v197 = v13 + int32(24)
						v204 = m.G0
						v206 = v204 - int32(16)
						m.G0 = v206
						v208 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
						if v208 <= int32(-4713) {
							if v208 != int32(-4713) {
								*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
								v285 = int32(-1)
							} else {
								v213 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
								if int32(10) < v213 {
									v224 = v213
									v225 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
									v226 = F_date2j(m, v208, v224, v225)
									mBase = m.M
									v229 = base.I64_extend_i32_s(v226 - int32(_a_F_parse_datetime_0))
									v230 = int64(63)
									F___multi3(m, v206, v229, v229>>(uint(v230)%64), int64(86400000000), int64(0))
									mBase = m.M
									v235 = *(*int64)(unsafe.Add(mBase, uint32(v206)+8))
									v236 = *(*int64)(unsafe.Add(mBase, uint32(v206)))
									if v235 != v236>>(uint(v230)%64) {
										*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
										v285 = int32(-1)
									} else {
										v241 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
										v242 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
										v243 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
										v244 = int32(60)
										v253 = base.I64_extend_i32_s(v194) + base.I64_extend_i32_s(v241+(v242+v243*v244)*v244)*int64(1000000)
										v254 = v236 + v253
										*(*int64)(unsafe.Add(mBase, uint32(v197))) = v254
										if base.B2i32(v253 < int64(0))^base.B2i32(v254 < v236) != 0 {
											*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
											v285 = int32(-1)
										} else {
											if base.Ui64(v254+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v285 = int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
												v285 = int32(-1)
											}
										}
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
									v285 = int32(-1)
								}
							}
						} else {
							if v208 <= int32(_a_F_parse_datetime_1) {
								v218 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
								v224 = v218
								v225 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
								v226 = F_date2j(m, v208, v224, v225)
								mBase = m.M
								v229 = base.I64_extend_i32_s(v226 - int32(_a_F_parse_datetime_0))
								v230 = int64(63)
								F___multi3(m, v206, v229, v229>>(uint(v230)%64), int64(86400000000), int64(0))
								mBase = m.M
								v235 = *(*int64)(unsafe.Add(mBase, uint32(v206)+8))
								v236 = *(*int64)(unsafe.Add(mBase, uint32(v206)))
								if v235 != v236>>(uint(v230)%64) {
									*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
									v285 = int32(-1)
								} else {
									v241 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
									v242 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
									v243 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
									v244 = int32(60)
									v253 = base.I64_extend_i32_s(v194) + base.I64_extend_i32_s(v241+(v242+v243*v244)*v244)*int64(1000000)
									v254 = v236 + v253
									*(*int64)(unsafe.Add(mBase, uint32(v197))) = v254
									if base.B2i32(v253 < int64(0))^base.B2i32(v254 < v236) != 0 {
										*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
										v285 = int32(-1)
									} else {
										if base.Ui64(v254+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
											v285 = int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
											v285 = int32(-1)
										}
									}
								}
							} else {
								if v208 != int32(_a_F_parse_datetime_2) {
									*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
									v285 = int32(-1)
								} else {
									v221 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
									if int32(5) < v221 {
										*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
										v285 = int32(-1)
									} else {
										v224 = v221
										v225 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
										v226 = F_date2j(m, v208, v224, v225)
										mBase = m.M
										v229 = base.I64_extend_i32_s(v226 - int32(_a_F_parse_datetime_0))
										v230 = int64(63)
										F___multi3(m, v206, v229, v229>>(uint(v230)%64), int64(86400000000), int64(0))
										mBase = m.M
										v235 = *(*int64)(unsafe.Add(mBase, uint32(v206)+8))
										v236 = *(*int64)(unsafe.Add(mBase, uint32(v206)))
										if v235 != v236>>(uint(v230)%64) {
											*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
											v285 = int32(-1)
										} else {
											v241 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
											v242 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
											v243 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
											v244 = int32(60)
											v253 = base.I64_extend_i32_s(v194) + base.I64_extend_i32_s(v241+(v242+v243*v244)*v244)*int64(1000000)
											v254 = v236 + v253
											*(*int64)(unsafe.Add(mBase, uint32(v197))) = v254
											if base.B2i32(v253 < int64(0))^base.B2i32(v254 < v236) != 0 {
												*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
												v285 = int32(-1)
											} else {
												if base.Ui64(v254+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
													v285 = int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(0)
													v285 = int32(-1)
												}
											}
										}
									}
								}
							}
						}
						m.G0 = v206 + int32(16)
						if v285 != 0 {
							v289 = F_errsave_start(m, l5)
							mBase = m.M
							v290 = m.ExcPending
							if v290 != 0 {
								return int32(0)
							} else {
								if v289 == int32(0) {
									v581 = v191
									m.G0 = v13 + int32(96)
									return v581
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v295 = m.ExcPending
									if v295 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_parse_datetime_9), int32(0))
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_10), int32(_a_F_parse_datetime_6))
											mBase = m.M
											v304 = m.ExcPending
											if v304 != 0 {
												return int32(0)
											} else {
												v581 = v191
												m.G0 = v13 + int32(96)
												return v581
											}
										}
									}
								}
							}
						} else {
							v307 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							F_AdjustTimestampForTypmod(m, v13+int32(24), v307, l5)
							mBase = m.M
							v309 = m.ExcPending
							if v309 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1114)
								v312 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
								v313 = F_Int64GetDatum(m, v312)
								mBase = m.M
								v314 = m.ExcPending
								if v314 != 0 {
									return int32(0)
								} else {
									v581 = v313
									m.G0 = v13 + int32(96)
									return v581
								}
							}
						}
					}
				} else {
					if v43 != 0 {
						v315 = int32(0)
						v316 = F_errsave_start(m, l5)
						mBase = m.M
						v317 = m.ExcPending
						if v317 != 0 {
							return int32(0)
						} else {
							if v316 == int32(0) {
								v581 = v315
								m.G0 = v13 + int32(96)
								return v581
							} else {
								F_errcode(m, int32(117440642))
								mBase = m.M
								v322 = m.ExcPending
								if v322 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_parse_datetime_11), int32(0))
									mBase = m.M
									v326 = m.ExcPending
									if v326 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_12), int32(_a_F_parse_datetime_6))
										mBase = m.M
										v331 = m.ExcPending
										if v331 != 0 {
											return int32(0)
										} else {
											v581 = v315
											m.G0 = v13 + int32(96)
											return v581
										}
									}
								}
							}
						}
					} else {
						v332 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
						if v332 <= int32(-4713) {
							if v332 != int32(-4713) {
								v349 = int32(0)
								v350 = F_errsave_start(m, l5)
								mBase = m.M
								v351 = m.ExcPending
								if v351 != 0 {
									return int32(0)
								} else {
									if v350 == int32(0) {
										v581 = v349
										m.G0 = v13 + int32(96)
										return v581
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v356 = m.ExcPending
										if v356 != 0 {
											return int32(0)
										} else {
											v357 = F_text_to_cstring(m, l0)
											mBase = m.M
											v358 = m.ExcPending
											if v358 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v357
												F_errmsg(m, int32(_a_F_parse_datetime_13), v13+int32(16))
												mBase = m.M
												v364 = m.ExcPending
												if v364 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
													mBase = m.M
													v369 = m.ExcPending
													if v369 != 0 {
														return int32(0)
													} else {
														v581 = v349
														m.G0 = v13 + int32(96)
														return v581
													}
												}
											}
										}
									}
								}
							} else {
								v337 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
								if v337 <= int32(10) {
									v349 = int32(0)
									v350 = F_errsave_start(m, l5)
									mBase = m.M
									v351 = m.ExcPending
									if v351 != 0 {
										return int32(0)
									} else {
										if v350 == int32(0) {
											v581 = v349
											m.G0 = v13 + int32(96)
											return v581
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v356 = m.ExcPending
											if v356 != 0 {
												return int32(0)
											} else {
												v357 = F_text_to_cstring(m, l0)
												mBase = m.M
												v358 = m.ExcPending
												if v358 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v357
													F_errmsg(m, int32(_a_F_parse_datetime_13), v13+int32(16))
													mBase = m.M
													v364 = m.ExcPending
													if v364 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
														mBase = m.M
														v369 = m.ExcPending
														if v369 != 0 {
															return int32(0)
														} else {
															v581 = v349
															m.G0 = v13 + int32(96)
															return v581
														}
													}
												}
											}
										}
									}
								} else {
									v521 = v337
									v522 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
									v527 = base.B2i32(int32(2) < v521)
									if int32(2) < v521 {
										v528 = int32(_a_F_parse_datetime_15)
									} else {
										v528 = int32(_a_F_parse_datetime_16)
									}
									v529 = v528 + v332
									v534 = base.I32_div_s(v529, int32(4))
									v537 = base.I32_div_s(v529, int32(-100))
									v540 = base.I32_div_s(v529, int32(400))
									if int32(2) < v521 {
										v544 = int32(1)
									} else {
										v544 = int32(13)
									}
									v549 = base.I32_div_s((v544+v521)*int32(_a_F_parse_datetime_17), int32(256))
									v552 = v522 + v529*int32(365) + v534 + v537 + v540 + v549 - int32(_a_F_parse_datetime_18)
									if base.Ui32(int32(2147483494)) <= base.Ui32(v552) {
										v555 = int32(0)
										v556 = F_errsave_start(m, l5)
										mBase = m.M
										v557 = m.ExcPending
										if v557 != 0 {
											return int32(0)
										} else {
											if v556 == int32(0) {
												v581 = v555
												m.G0 = v13 + int32(96)
												return v581
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v562 = m.ExcPending
												if v562 != 0 {
													return int32(0)
												} else {
													v563 = F_text_to_cstring(m, l0)
													mBase = m.M
													v564 = m.ExcPending
													if v564 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13))) = v563
														F_errmsg(m, int32(_a_F_parse_datetime_13), v13)
														mBase = m.M
														v568 = m.ExcPending
														if v568 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_19), int32(_a_F_parse_datetime_6))
															mBase = m.M
															v573 = m.ExcPending
															if v573 != 0 {
																return int32(0)
															} else {
																v581 = v555
																m.G0 = v13 + int32(96)
																return v581
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1082)
										v581 = v552 - int32(_a_F_parse_datetime_0)
										m.G0 = v13 + int32(96)
										return v581
									}
								}
							}
						} else {
							if v332 <= int32(_a_F_parse_datetime_1) {
								v342 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
								v521 = v342
								v522 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
								v527 = base.B2i32(int32(2) < v521)
								if int32(2) < v521 {
									v528 = int32(_a_F_parse_datetime_15)
								} else {
									v528 = int32(_a_F_parse_datetime_16)
								}
								v529 = v528 + v332
								v534 = base.I32_div_s(v529, int32(4))
								v537 = base.I32_div_s(v529, int32(-100))
								v540 = base.I32_div_s(v529, int32(400))
								if int32(2) < v521 {
									v544 = int32(1)
								} else {
									v544 = int32(13)
								}
								v549 = base.I32_div_s((v544+v521)*int32(_a_F_parse_datetime_17), int32(256))
								v552 = v522 + v529*int32(365) + v534 + v537 + v540 + v549 - int32(_a_F_parse_datetime_18)
								if base.Ui32(int32(2147483494)) <= base.Ui32(v552) {
									v555 = int32(0)
									v556 = F_errsave_start(m, l5)
									mBase = m.M
									v557 = m.ExcPending
									if v557 != 0 {
										return int32(0)
									} else {
										if v556 == int32(0) {
											v581 = v555
											m.G0 = v13 + int32(96)
											return v581
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v562 = m.ExcPending
											if v562 != 0 {
												return int32(0)
											} else {
												v563 = F_text_to_cstring(m, l0)
												mBase = m.M
												v564 = m.ExcPending
												if v564 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13))) = v563
													F_errmsg(m, int32(_a_F_parse_datetime_13), v13)
													mBase = m.M
													v568 = m.ExcPending
													if v568 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_19), int32(_a_F_parse_datetime_6))
														mBase = m.M
														v573 = m.ExcPending
														if v573 != 0 {
															return int32(0)
														} else {
															v581 = v555
															m.G0 = v13 + int32(96)
															return v581
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1082)
									v581 = v552 - int32(_a_F_parse_datetime_0)
									m.G0 = v13 + int32(96)
									return v581
								}
							} else {
								if v332 != int32(_a_F_parse_datetime_2) {
									v349 = int32(0)
									v350 = F_errsave_start(m, l5)
									mBase = m.M
									v351 = m.ExcPending
									if v351 != 0 {
										return int32(0)
									} else {
										if v350 == int32(0) {
											v581 = v349
											m.G0 = v13 + int32(96)
											return v581
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v356 = m.ExcPending
											if v356 != 0 {
												return int32(0)
											} else {
												v357 = F_text_to_cstring(m, l0)
												mBase = m.M
												v358 = m.ExcPending
												if v358 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v357
													F_errmsg(m, int32(_a_F_parse_datetime_13), v13+int32(16))
													mBase = m.M
													v364 = m.ExcPending
													if v364 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
														mBase = m.M
														v369 = m.ExcPending
														if v369 != 0 {
															return int32(0)
														} else {
															v581 = v349
															m.G0 = v13 + int32(96)
															return v581
														}
													}
												}
											}
										}
									}
								} else {
									v345 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
									if v345 < int32(6) {
										v521 = v345
										v522 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
										v527 = base.B2i32(int32(2) < v521)
										if int32(2) < v521 {
											v528 = int32(_a_F_parse_datetime_15)
										} else {
											v528 = int32(_a_F_parse_datetime_16)
										}
										v529 = v528 + v332
										v534 = base.I32_div_s(v529, int32(4))
										v537 = base.I32_div_s(v529, int32(-100))
										v540 = base.I32_div_s(v529, int32(400))
										if int32(2) < v521 {
											v544 = int32(1)
										} else {
											v544 = int32(13)
										}
										v549 = base.I32_div_s((v544+v521)*int32(_a_F_parse_datetime_17), int32(256))
										v552 = v522 + v529*int32(365) + v534 + v537 + v540 + v549 - int32(_a_F_parse_datetime_18)
										if base.Ui32(int32(2147483494)) <= base.Ui32(v552) {
											v555 = int32(0)
											v556 = F_errsave_start(m, l5)
											mBase = m.M
											v557 = m.ExcPending
											if v557 != 0 {
												return int32(0)
											} else {
												if v556 == int32(0) {
													v581 = v555
													m.G0 = v13 + int32(96)
													return v581
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v562 = m.ExcPending
													if v562 != 0 {
														return int32(0)
													} else {
														v563 = F_text_to_cstring(m, l0)
														mBase = m.M
														v564 = m.ExcPending
														if v564 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13))) = v563
															F_errmsg(m, int32(_a_F_parse_datetime_13), v13)
															mBase = m.M
															v568 = m.ExcPending
															if v568 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_19), int32(_a_F_parse_datetime_6))
																mBase = m.M
																v573 = m.ExcPending
																if v573 != 0 {
																	return int32(0)
																} else {
																	v581 = v555
																	m.G0 = v13 + int32(96)
																	return v581
																}
															}
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1082)
											v581 = v552 - int32(_a_F_parse_datetime_0)
											m.G0 = v13 + int32(96)
											return v581
										}
									} else {
										v349 = int32(0)
										v350 = F_errsave_start(m, l5)
										mBase = m.M
										v351 = m.ExcPending
										if v351 != 0 {
											return int32(0)
										} else {
											if v350 == int32(0) {
												v581 = v349
												m.G0 = v13 + int32(96)
												return v581
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return int32(0)
												} else {
													v357 = F_text_to_cstring(m, l0)
													mBase = m.M
													v358 = m.ExcPending
													if v358 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v357
														F_errmsg(m, int32(_a_F_parse_datetime_13), v13+int32(16))
														mBase = m.M
														v364 = m.ExcPending
														if v364 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
															mBase = m.M
															v369 = m.ExcPending
															if v369 != 0 {
																return int32(0)
															} else {
																v581 = v349
																m.G0 = v13 + int32(96)
																return v581
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
			} else {
				if v39 != 0 {
					if v37&int32(4) != 0 {
						v373 = F_palloc(m, int32(16))
						mBase = m.M
						v374 = m.ExcPending
						if v374 != 0 {
							return int32(0)
						} else {
							v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+44)))
							if v375 == int32(1) {
								v378 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v378
								v380 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
								v382 = v13 + int32(52)
								v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
								v384 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
								v385 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v373)+8)) = v378
								v388 = int32(60)
								*(*int64)(unsafe.Add(mBase, uint32(v373))) = base.I64_extend_i32_s(v380) + base.I64_extend_i32_s(v383+(v384+v385*v388)*v388)*int64(1000000)
								v416 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								if base.Ui32(v416) <= base.Ui32(int32(6)) {
									v423 = v416 << (uint(int32(3)) % 32)
									v426 = *(*int64)(unsafe.Add(mBase, uint32(v423)+uint32(_c_F_parse_datetime[0])))
									v429 = *(*int64)(unsafe.Add(mBase, uint32(v423)+uint32(_c_F_parse_datetime[1])))
									v430 = *(*int64)(unsafe.Add(mBase, uint32(v373)))
									if int64(0) <= v430 {
										v433 = v429 + v430
										v434 = base.I64_rem_s(v433, v426)
										v440 = v433 - v434
									} else {
										v436 = v429 - v430
										v437 = base.I64_rem_s(v436, v426)
										v440 = v437 - v436
									}
									*(*int64)(unsafe.Add(mBase, uint32(v373))) = v440
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1266)
								v581 = v373
								m.G0 = v13 + int32(96)
								return v581
							} else {
								v399 = int32(0)
								v400 = F_errsave_start(m, l5)
								mBase = m.M
								v401 = m.ExcPending
								if v401 != 0 {
									return int32(0)
								} else {
									if v400 == int32(0) {
										v581 = v399
										m.G0 = v13 + int32(96)
										return v581
									} else {
										F_errcode(m, int32(117440642))
										mBase = m.M
										v406 = m.ExcPending
										if v406 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_parse_datetime_20), int32(0))
											mBase = m.M
											v410 = m.ExcPending
											if v410 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_21), int32(_a_F_parse_datetime_6))
												mBase = m.M
												v415 = m.ExcPending
												if v415 != 0 {
													return int32(0)
												} else {
													v581 = v399
													m.G0 = v13 + int32(96)
													return v581
												}
											}
										}
									}
								}
							}
						}
					} else {
						v449 = v13 + int32(24)
						v450 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+40)))
						v452 = v13 + int32(52)
						v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
						v454 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
						v455 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
						v456 = int32(60)
						*(*int64)(unsafe.Add(mBase, uint32(v449))) = v450 + base.I64_extend_i32_s(v453+(v454+v455*v456)*v456)*int64(1000000)
						v469 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						if base.Ui32(v469) <= base.Ui32(int32(6)) {
							v476 = v469 << (uint(int32(3)) % 32)
							v479 = *(*int64)(unsafe.Add(mBase, uint32(v476)+uint32(_c_F_parse_datetime[0])))
							v482 = *(*int64)(unsafe.Add(mBase, uint32(v476)+uint32(_c_F_parse_datetime[1])))
							v483 = *(*int64)(unsafe.Add(mBase, uint32(v449)))
							if int64(0) <= v483 {
								v486 = v482 + v483
								v487 = base.I64_rem_s(v486, v479)
								v493 = v486 - v487
							} else {
								v489 = v482 - v483
								v490 = base.I64_rem_s(v489, v479)
								v493 = v490 - v489
							}
							*(*int64)(unsafe.Add(mBase, uint32(v449))) = v493
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1083)
						v501 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
						v502 = F_Int64GetDatum(m, v501)
						mBase = m.M
						v503 = m.ExcPending
						if v503 != 0 {
							return int32(0)
						} else {
							v581 = v502
							m.G0 = v13 + int32(96)
							return v581
						}
					}
				} else {
					v504 = int32(0)
					v505 = F_errsave_start(m, l5)
					mBase = m.M
					v506 = m.ExcPending
					if v506 != 0 {
						return int32(0)
					} else {
						if v505 == int32(0) {
							v581 = v504
							m.G0 = v13 + int32(96)
							return v581
						} else {
							F_errcode(m, int32(117440642))
							mBase = m.M
							v511 = m.ExcPending
							if v511 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_parse_datetime_22), int32(0))
								mBase = m.M
								v515 = m.ExcPending
								if v515 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_23), int32(_a_F_parse_datetime_6))
									mBase = m.M
									v520 = m.ExcPending
									if v520 != 0 {
										return int32(0)
									} else {
										v581 = v504
										m.G0 = v13 + int32(96)
										return v581
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
func F_parse_ident_line(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v21 = v19
	goto L3
L2:
	;
	v21 = int32(0)
	goto L3
L3:
	;
	v22 = int32(16)
	v23 = l0 + v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = F_palloc0(m, v22)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v17
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if int32(2) <= v32 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	m.G0 = v15 + int32(80)
	return v378
L7:
	;
	v35 = int32(0)
	v37 = F_errstart(m, l1, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = F_pstrdup(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L20
	}
L10:
	;
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v60 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L19
	}
L14:
	;
	F_errmsg(m, int32(_a_F_parse_ident_line_0), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2769), int32(_a_F_parse_ident_line_3))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v60
	v378 = v35
	goto L6
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v66
	v70 = v21 + int32(4)
	if v70 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if int32(2) <= v110 {
		goto L36
	} else {
		goto L37
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if base.Ui32(v70) < base.Ui32(v72+v73<<(uint(int32(2))%32)) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v79 = int32(0)
	v81 = F_errstart(m, l1, v79)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	if v81 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v106 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L35
	}
L30:
	;
	F_errmsg(m, int32(_a_F_parse_ident_line_4), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v17
	F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2775), int32(_a_F_parse_ident_line_3))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v106
	v378 = v79
	goto L6
L36:
	;
	v113 = int32(0)
	v115 = F_errstart(m, l1, v113)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+4)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if v146&int32(3) == int32(0) {
		v170 = v146
		goto L51
	} else {
		goto L52
	}
L39:
	;
	if v115 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v140 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L48
	}
L43:
	;
	F_errmsg(m, int32(_a_F_parse_ident_line_0), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v17
	F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15+int32(32))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2777), int32(_a_F_parse_ident_line_3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v140
	v378 = v113
	goto L6
L49:
	;
	v206 = F_palloc0(m, v203+int32(13))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L66
	}
L50:
	;
	v203 = v195 - v146
	goto L49
L51:
	;
	v174 = v170
	goto L60
L52:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v154 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v203 = int32(0)
	goto L49
L54:
	;
	goto L55
L55:
	;
	v159 = v146
	goto L56
L56:
	;
	v163 = v159 + int32(1)
	if v163&int32(3) == int32(0) {
		v170 = v163
		goto L51
	} else {
		goto L58
	}
L57:
	;
	v195 = v163
	goto L50
L58:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v168 != 0 {
		v159 = v163
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v183 = int32(-2139062144)
	if (int32(16843008)-v180|v180)&v183 == v183 {
		v174 = v174 + int32(4)
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v189 = v174
	goto L63
L62:
	;
	goto L61
L63:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v193 != 0 {
		v189 = v189 + int32(1)
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v195 = v189
	goto L50
L65:
	;
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+8)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)) = uint8(v145)
	v212 = v206 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v212
	v215 = v203 + int32(1)
	if v215 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v206
	v220 = v21 + int32(8)
	if v220 != 0 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v216 = F__emscripten_memcpy_bulkmem(m, v212, v146, v215)
	mBase = m.M
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if int32(2) <= v260 {
		goto L86
	} else {
		goto L87
	}
L72:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if base.Ui32(v220) < base.Ui32(v222+v223<<(uint(int32(2))%32)) {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v230 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	if v230 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v255 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L85
	}
L80:
	;
	F_errmsg(m, int32(_a_F_parse_ident_line_4), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v17
	F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15+int32(48))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2785), int32(_a_F_parse_ident_line_3))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	goto L79
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v255
	v378 = int32(0)
	goto L6
L86:
	;
	v263 = int32(0)
	v265 = F_errstart(m, l1, v263)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+4)))
	v296 = int32(0)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	if v297&int32(3) == v296 {
		v321 = v297
		goto L101
	} else {
		goto L102
	}
L89:
	;
	if v265 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v290 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L98
	}
L93:
	;
	F_errmsg(m, int32(_a_F_parse_ident_line_0), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v17
	F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15-int32(-64))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2787), int32(_a_F_parse_ident_line_3))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	goto L92
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v290
	v378 = v263
	goto L6
L99:
	;
	v357 = F_palloc0(m, v354+int32(13))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L116
	}
L100:
	;
	v354 = v346 - v297
	goto L99
L101:
	;
	v325 = v321
	goto L110
L102:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	if v305 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v354 = int32(0)
	goto L99
L104:
	;
	goto L105
L105:
	;
	v310 = v297
	goto L106
L106:
	;
	v314 = v310 + int32(1)
	if v314&int32(3) == int32(0) {
		v321 = v314
		goto L101
	} else {
		goto L108
	}
L107:
	;
	v346 = v314
	goto L100
L108:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v319 != 0 {
		v310 = v314
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v334 = int32(-2139062144)
	if (int32(16843008)-v331|v331)&v334 == v334 {
		v325 = v325 + int32(4)
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v340 = v325
	goto L113
L112:
	;
	goto L111
L113:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v344 != 0 {
		v340 = v340 + int32(1)
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v346 = v340
	goto L100
L115:
	;
	goto L114
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+8)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v357)+4)) = uint8(v295)
	v363 = v357 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v363
	v366 = v354 + int32(1)
	if v366 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v357
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v371 = F_regcomp_auth_token(m, v370, v24, v17, v23, l1)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L121
	}
L118:
	;
	v367 = F__emscripten_memcpy_bulkmem(m, v363, v297, v366)
	mBase = m.M
	goto L120
L119:
	;
	goto L120
L120:
	;
	goto L117
L121:
	;
	if v371 != 0 {
		v378 = v296
		goto L6
	} else {
		goto L122
	}
L122:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v375 = F_regcomp_auth_token(m, v374, v24, v17, v23, l1)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	if v375 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v377 = int32(0)
	goto L126
L125:
	;
	v377 = v26
	goto L126
L126:
	;
	v378 = v377
	goto L6
}
func F_pathkeys_count_contained_in(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	if l0 == l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v76
	return int32(1)
L2:
	;
	if l0 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	return int32(1)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	return int32(1)
L7:
	;
	goto L8
L8:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23
	return v23
L10:
	;
	goto L11
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = int32(0)
	if v28 < v27 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = v27
	goto L14
L13:
	;
	v31 = v28
	goto L14
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = int32(0)
	goto L15
L15:
	;
	if v36 < v32 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = v44 + v36<<(uint(int32(2))%32)
	goto L19
L18:
	;
	v48 = int32(0)
	goto L19
L19:
	;
	if v36 == v31 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31
	return base.B2i32(v48 == int32(0))
L21:
	;
	goto L22
L22:
	;
	v55 = base.B2i32(v48 == int32(0))
	if v48 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
	return v55
L24:
	;
	goto L25
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v63 = v60 + v36<<(uint(int32(2))%32)
	if v63 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
	return v55
L27:
	;
	goto L28
L28:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v68 != v69 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
	return int32(0)
L30:
	;
	v36 = v36 + int32(1)
	goto L15
}
func F_perform_work_item(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v270 int64
	_ = v270
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v414 int32
	_ = v414
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v467 int32
	_ = v467
	var v479 int32
	_ = v479
	var v491 int32
	_ = v491
	var v503 int32
	_ = v503
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v630 int32
	_ = v630
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int64
	_ = v648
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(96)
	m.G0 = v23
	v26 = l0 + int32(12)
	v28 = l0 + int32(16)
	v32 = v2
	v33 = v2
	v34 = v2
	v35 = v2
	v36 = v2
	v37 = v2
	v38 = v2
	v39 = v2
	v40 = v2
	v41 = v2
	v42 = int32(-1)
	v43 = v2
	v47 = v23
	goto L1
L1:
	;
	goto L4
L2:
	;
	m.G0 = v23 + int32(96)
	return
L3:
	;
	goto L2
L4:
	;
	if v42 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	goto L3
L6:
	;
	v647 = int32(m.ExcTag)
	v648 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v647 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L7:
	;
	v600 = v591 & int32(1)
	if v600 != 0 {
		goto L79
	} else {
		goto L80
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v555
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v556
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v554
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v553
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v551
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v552
	v575 = int32(1)
	v576 = v559 & v575
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v576)
	v579 = v560 & v575
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v579)
	F_pfree(m, v554)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		v644 = v564
		goto L6
	} else {
		goto L78
	}
L9:
	;
	if v115 == int32(0) {
		v583 = v76
		v584 = v59
		v585 = v101
		v586 = v115
		v587 = v36
		v588 = v37
		v589 = v38
		v590 = v26
		v591 = v118
		v592 = v120
		v596 = v59
		goto L7
	} else {
		goto L77
	}
L10:
	;
	v53 = v47 - int32(208)
	m.G0 = v53
	v56 = v53 - int32(16)
	m.G0 = v56
	v59 = v56 - int32(160)
	m.G0 = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v38
	v65 = int32(1)
	v66 = v40 & v65
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v66)
	v69 = v41 & v65
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	v76 = F_get_rel_name(m, v61)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		v644 = v59
		goto L6
	} else {
		goto L13
	}
L11:
	;
	v308 = v32
	v309 = v33
	v310 = v34
	v311 = v35
	v312 = v36
	v313 = v37
	v314 = v38
	v315 = v39
	v316 = v40
	v317 = v41
	v319 = v43
	v321 = v47
	goto L12
L12:
	;
	if v319 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v38
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v66)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	v89 = F_get_rel_namespace(m, v78)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		v644 = v59
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v66)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v69)
	v101 = F_get_namespace_name(m, v89)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		v644 = v59
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v66)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v69)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[0]))
	v115 = F_get_database_name(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		v644 = v59
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v117 = int32(0)
	v118 = base.B2i32(v101 != v117)
	v120 = base.B2i32(v76 != v117)
	if v76 == v117 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	if v101 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	if v115 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v127 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v38
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v118)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v120)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	v143 = F_pg_snprintf(m, v53, int32(184), int32(_a_F_perform_work_item_0), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		v644 = v59
		goto L6
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v38
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v118)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v120)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	if v53&int32(3) == int32(0) {
		v178 = v53
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L22
L24:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v212 != int32(-1) {
		goto L42
	} else {
		goto L43
	}
L25:
	;
	v211 = v203 - v53
	goto L24
L26:
	;
	v182 = v178
	goto L35
L27:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v162 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v211 = int32(0)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v167 = v53
	goto L31
L31:
	;
	v171 = v167 + int32(1)
	if v171&int32(3) == int32(0) {
		v178 = v171
		goto L26
	} else {
		goto L33
	}
L32:
	;
	v203 = v171
	goto L25
L33:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v176 != 0 {
		v167 = v171
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v191 = int32(-2139062144)
	if (int32(16843008)-v188|v188)&v191 == v191 {
		v182 = v182 + int32(4)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v197 = v182
	goto L38
L37:
	;
	goto L36
L38:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v201 != 0 {
		v197 = v197 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v203 = v197
	goto L25
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v28
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v118)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v120)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v101
	v253 = F_pg_snprintf(m, v211+v53, int32(184)-v211, int32(_a_F_perform_work_item_1), v23+int32(32))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		v644 = v59
		goto L6
	} else {
		goto L46
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v28
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v118)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v120)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v212
	v230 = F_pg_snprintf(m, v56, int32(14), int32(_a_F_perform_work_item_2), v23+int32(48))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		v644 = v59
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v232)
	goto L41
L45:
	;
	goto L41
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v28
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v118)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v120)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[1]))
	if v266 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v28
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v118)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v120)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	F_pgstat_report_activity(m, int32(3), v53)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v28
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v118)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v120)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v59
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[2]))
	F_MemoryContextReset(m, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		v644 = v59
		goto L6
	} else {
		goto L51
	}
L48:
	;
	v270 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_perform_work_item[3])) = v270
	goto L50
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[4]))
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[5]))
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v23 + int32(56)
	goto L55
L53:
	;
	v308 = v76
	v309 = v59
	v310 = v101
	v311 = v115
	v312 = v302
	v313 = v300
	v314 = v28
	v315 = v26
	v316 = v118
	v317 = v120
	v319 = int32(0)
	v321 = v59
	goto L12
L55:
	;
	goto L53
L56:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[4])) = v313
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[5])) = v312
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[7])) = v547
	v551 = v308
	v552 = v309
	v553 = v310
	v554 = v311
	v555 = v312
	v556 = v313
	v557 = v314
	v558 = v315
	v559 = v316
	v560 = v317
	v564 = v321
	goto L8
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[5])) = v309
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[7])) = v328
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v330 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[4])) = v313
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[5])) = v312
	v427 = int32(_a_F_perform_work_item_3)
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[8]))
	v430 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[8])) = v429 + v430
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	v442 = v316 & v430
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v442)
	v445 = v317 & v430
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v445)
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L70
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[9])) = int32(0)
	goto L56
L61:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	v334 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v314))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	v338 = int32(1)
	v339 = v316 & v338
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v339)
	v342 = v317 & v338
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v342)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	v349 = F_Int64GetDatum(m, v334)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	v373 = int32(1)
	v374 = v316 & v373
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v374)
	v377 = v317 & v373
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v377)
	v381 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L66
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v339)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v342)
	v363 = F_DirectFunctionCall2Coll(m, int32(16), int32(0), v333, v349)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L65
	}
L65:
	;
	goto L60
L66:
	;
	if v381 == int32(0) {
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v374)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v377)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v385
	F_errmsg_internal(m, int32(_a_F_perform_work_item_4), v23)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v374)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v377)
	F_errfinish(m, int32(_a_F_perform_work_item_5), int32(2658), int32(_a_F_perform_work_item_6))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L69
	}
L69:
	;
	goto L60
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v442)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v445)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v311
	F_errcontext_msg(m, int32(_a_F_perform_work_item_7), v23+int32(16))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v442)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v445)
	F_EmitErrorReport(m)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v442)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v445)
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v442)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v445)
	F_FlushErrorState(m)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v442)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v445)
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[2]))
	F_MemoryContextReset(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v309
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v442)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v445)
	F_StartTransactionCommand(m)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		v644 = v321
		goto L6
	} else {
		goto L76
	}
L76:
	;
	v530 = int32(_a_F_perform_work_item_3)
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_perform_work_item[8])) = v532 - int32(1)
	goto L56
L77:
	;
	v551 = v76
	v552 = v59
	v553 = v101
	v554 = v115
	v555 = v36
	v556 = v37
	v557 = v38
	v558 = v26
	v559 = v118
	v560 = v120
	v564 = v59
	goto L8
L78:
	;
	v583 = v551
	v584 = v552
	v585 = v553
	v586 = v554
	v587 = v555
	v588 = v556
	v589 = v557
	v590 = v558
	v591 = v559
	v592 = v560
	v596 = v564
	goto L7
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v590
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v584
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v600)
	v611 = v592 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v611)
	F_pfree(m, v585)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		v644 = v596
		goto L6
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v616 = v592 & int32(1)
	if v616 == int32(0) {
		goto L3
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v590
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v584
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)) = uint8(v600)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)) = uint8(v616)
	F_pfree(m, v583)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		v644 = v596
		goto L6
	} else {
		goto L84
	}
L84:
	;
	goto L5
L85:
	;
	v652 = int32(v648)
	m.G0 = v644
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v652)))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	if v23+int32(56) == v659 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	m.ExcPending = 1
	goto L94
L87:
	;
	if v662 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	v662 = v661
	goto L90
L89:
	;
	v662 = int32(0)
	goto L90
L90:
	;
	goto L87
L91:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+75)))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+74)))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v32 = v665
	v33 = v663
	v34 = v666
	v35 = v667
	v36 = v671
	v37 = v672
	v38 = v670
	v39 = v664
	v40 = v669
	v41 = v668
	v42 = v662
	v43 = v654
	v47 = v644
	goto L1
L92:
	;
	goto L93
L93:
	;
	F___wasm_longjmp(m, v655, v654)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	return
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgarch_waken_stop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[1]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_pglz_decompress_datum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_palloc(m, v3&int32(1073741823)+int32(4))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = int32(8)
	v13 = l0 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = v8 + int32(4)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = v20 + v21&int32(1073741823)
	v33 = v13 + (int32(base.Ui32(v14)>>(uint(int32(2))%32)) - v12)
	if base.Ui32(v33) <= base.Ui32(v13) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v204 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L4:
	;
	goto L3
L5:
	;
	goto L41
L6:
	;
	v176 = v13
	v177 = v20
	goto L5
L7:
	;
	if base.Ui32(v32) <= base.Ui32(v20) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v36 = v13
	v37 = v20
	goto L9
L9:
	;
	v49 = v36 + int32(1)
	if base.Ui32(v33) <= base.Ui32(v49) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v176 = v162
	v177 = v163
	goto L5
L11:
	;
	if base.Ui32(v33) <= base.Ui32(v162) {
		v176 = v162
		v177 = v163
		goto L5
	} else {
		goto L39
	}
L12:
	;
	v162 = v49
	v163 = v37
	goto L11
L13:
	;
	goto L14
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v54 = v37
	v56 = v49
	v62 = v51
	v63 = int32(0)
	goto L15
L15:
	;
	if v62&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v162 = v139
	v163 = v151
	goto L11
L17:
	;
	if base.Ui32(int32(6)) < base.Ui32(v63) {
		v162 = v139
		v163 = v151
		goto L11
	} else {
		goto L36
	}
L18:
	;
	v67 = int32(-1)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v72 = v68&int32(15) + int32(3)
	if v72 != int32(18) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v133)
	v135 = int32(1)
	v139 = v56 + v135
	v151 = v54 + v135
	goto L17
L21:
	;
	v82 = v72
	v83 = v56 + int32(2)
	goto L23
L22:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+2)))
	v82 = v77 + int32(18)
	v83 = v56 + int32(3)
	goto L23
L23:
	;
	if base.Ui32(v33) < base.Ui32(v83) {
		v204 = v67
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v90 = v85 | v68<<(uint(int32(4))%32)&int32(3840)
	if v90 == int32(0) {
		v204 = v67
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v54-v20 < v90 {
		v204 = v67
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v95 = v32 - v54
	if v82 < v95 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v97 = v82
	goto L29
L28:
	;
	v97 = v95
	goto L29
L29:
	;
	if v90 < v97 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v100 = v54
	v102 = v90
	v104 = v97
	goto L33
L31:
	;
	v119 = v54
	v121 = v90
	v123 = v97
	goto L32
L32:
	;
	v131 = F___memcpy(m, v119, v119-v121, v123)
	mBase = m.M
	v139 = v83
	v151 = v131 + v123
	goto L17
L33:
	;
	v111 = v104 - v102
	v113 = F___memcpy(m, v100, v100-v102, v102)
	mBase = m.M
	v114 = v113 + v102
	v116 = v102 << (uint(int32(1)) % 32)
	if v116 < v111 {
		v100 = v114
		v102 = v116
		v104 = v111
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v119 = v114
	v121 = v116
	v123 = v111
	goto L32
L35:
	;
	goto L34
L36:
	;
	if base.Ui32(v33) <= base.Ui32(v139) {
		v162 = v139
		v163 = v151
		goto L11
	} else {
		goto L37
	}
L37:
	;
	v155 = int32(1)
	if base.Ui32(v151) < base.Ui32(v32) {
		v54 = v151
		v56 = v139
		v62 = int32(base.Ui32(v62&int32(254)) >> (uint(v155) % 32))
		v63 = v63 + v155
		goto L15
	} else {
		goto L38
	}
L38:
	;
	goto L16
L39:
	;
	if base.Ui32(v163) < base.Ui32(v32) {
		v36 = v162
		v37 = v163
		goto L9
	} else {
		goto L40
	}
L40:
	;
	goto L10
L41:
	;
	v188 = int32(-1)
	if v176 != v33 {
		v204 = v188
		goto L4
	} else {
		goto L44
	}
L43:
	;
	v204 = v177 - v20
	goto L4
L44:
	;
	if v177 != v32 {
		v204 = v188
		goto L4
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v204<<(uint(int32(2))%32) + int32(16)
	return v8
L49:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errmsg_internal(m, int32(_a_F_pglz_decompress_datum_0), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_pglz_decompress_datum_1), int32(98), int32(_a_F_pglz_decompress_datum_2))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgmem_call_sighandler(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	m.T0[l0].(func(*base.Module, int32))(m, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_pgmem_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_init[0])) = int32(_a_F_pgmem_init_0)
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_init[1])) = int32(_a_F_pgmem_init_1)
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_init[2])) = int32(_a_F_pgmem_init_2)
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_init[3])) = int32(_a_F_pgmem_init_3)
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_init[4])) = int32(_a_F_pgmem_init_4)
	return
}
func F_points_box(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v20 float64
	_ = v20
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v31 float64
	_ = v31
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_palloc(m, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		if base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
			if base.F64_lt(v14, v20) != 0 {
				v28 = v20
				v29 = v7
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v20)&int64(9223372036854775807)) {
					v28 = v20
					v29 = v7
				} else {
					v28 = v14
					v29 = v8
				}
			}
		} else {
			v28 = v14
			v29 = v8
		}
		*(*float64)(unsafe.Add(mBase, uint32(v10))) = v28
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
		*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v31
		v34 = v8 + int32(8)
		v35 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		if base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v42 = v7 + int32(8)
			v43 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
			if base.F64_gt(v43, v35) != 0 {
				v52 = v42
				v53 = v43
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v43)&int64(9223372036854775807)) {
					v52 = v42
					v53 = v43
				} else {
					v52 = v34
					v53 = v35
				}
			}
		} else {
			v52 = v34
			v53 = v35
		}
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v53
		v55 = *(*float64)(unsafe.Add(mBase, uint32(v52)))
		*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v55
		return v10
	}
}
func F_pqinitmask(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	*(*int64)(unsafe.Add(mBase, _c_F_pqinitmask[0])) = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = int64(-15032385537)
	*(*int64)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = int64(-15032385537)
	v31 = int32(_a_F_pqinitmask_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v32 & base.I32_rotl(int32(-2), int32(4))
	v59 = int32(_a_F_pqinitmask_1)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v60 & base.I32_rotl(int32(-2), int32(4))
	v87 = int32(_a_F_pqinitmask_0)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v88 & base.I32_rotl(int32(-2), int32(5))
	v115 = int32(_a_F_pqinitmask_1)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v116 & base.I32_rotl(int32(-2), int32(5))
	v143 = int32(_a_F_pqinitmask_0)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v144 & base.I32_rotl(int32(-2), int32(3))
	v171 = int32(_a_F_pqinitmask_1)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v172 & base.I32_rotl(int32(-2), int32(3))
	v199 = int32(_a_F_pqinitmask_0)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v200 & base.I32_rotl(int32(-2), int32(7))
	v227 = int32(_a_F_pqinitmask_1)
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v228 & base.I32_rotl(int32(-2), int32(7))
	v255 = int32(_a_F_pqinitmask_0)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v256 & base.I32_rotl(int32(-2), int32(10))
	v283 = int32(_a_F_pqinitmask_1)
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v284 & base.I32_rotl(int32(-2), int32(10))
	v311 = int32(_a_F_pqinitmask_0)
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v312 & base.I32_rotl(int32(-2), int32(6))
	v339 = int32(_a_F_pqinitmask_1)
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v340 & base.I32_rotl(int32(-2), int32(6))
	v367 = int32(_a_F_pqinitmask_0)
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v368 & base.I32_rotl(int32(-2), int32(30))
	v395 = int32(_a_F_pqinitmask_1)
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v396 & base.I32_rotl(int32(-2), int32(30))
	v423 = int32(_a_F_pqinitmask_0)
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v424 & base.I32_rotl(int32(-2), int32(17))
	v451 = int32(_a_F_pqinitmask_1)
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v452 & base.I32_rotl(int32(-2), int32(17))
	v479 = int32(_a_F_pqinitmask_1)
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v480 & base.I32_rotl(int32(-2), int32(2))
	v507 = int32(_a_F_pqinitmask_1)
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v508 & base.I32_rotl(int32(-2), int32(14))
	v535 = int32(_a_F_pqinitmask_1)
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v536 & base.I32_rotl(int32(-2), int32(13))
	return
}
func F_pqsignal_be(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(144)
	m.G0 = v6
	switch l1 + int32(2) {
	case 0, 2:
		v16 = l1
	default:
		*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_pqsignal_be[0]))) = l1
		v16 = int32(_a_F_pqsignal_be_0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v6+int32(8)))) = int64(0)
	if l0 == int32(17) {
		v26 = int32(268435457)
	} else {
		v26 = int32(268435456)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v6)+136)) = v26
	v29 = v6 + int32(4)
	if base.Ui32(int32(65)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _c_F_pqsignal_be[1])) = int32(28)
	} else {
		if v29 != 0 {
			v35 = int32(140)
			v40 = F__emscripten_memcpy_bulkmem(m, l0*v35+int32(_a_F_pqsignal_be_1), v29, v35)
			mBase = m.M
		} else {
		}
	}
	m.G0 = v6 + int32(144)
	return
}
func F_pre_format_elog_string(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_pre_format_elog_string[0])) = l0
	return
}
func F_predicatelock_twophase_recover(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v13 {
	case 0:
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
		v19 = F_LWLockAcquire(m, v15+int32(3584), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			if v23 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v280 = m.ExcPending
				if v280 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_predicatelock_twophase_recover_0))
					mBase = m.M
					v283 = m.ExcPending
					if v283 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_predicatelock_twophase_recover_1), int32(0))
						mBase = m.M
						v287 = m.ExcPending
						if v287 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_predicatelock_twophase_recover_2), int32(_a_F_predicatelock_twophase_recover_3), int32(_a_F_predicatelock_twophase_recover_4))
							mBase = m.M
							v292 = m.ExcPending
							if v292 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if v23 == v22 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v280 = m.ExcPending
					if v280 != 0 {
						return
					} else {
						F_errcode(m, int32(_a_F_predicatelock_twophase_recover_0))
						mBase = m.M
						v283 = m.ExcPending
						if v283 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_predicatelock_twophase_recover_1), int32(0))
							mBase = m.M
							v287 = m.ExcPending
							if v287 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_predicatelock_twophase_recover_2), int32(_a_F_predicatelock_twophase_recover_3), int32(_a_F_predicatelock_twophase_recover_4))
								mBase = m.M
								v292 = m.ExcPending
								if v292 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
					v33 = v22 + int32(8)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
					if v34 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v33
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v33
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v40
					*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v33))) = v23
					v45 = v23 + int32(-64)
					if v45 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v280 = m.ExcPending
						if v280 != 0 {
							return
						} else {
							F_errcode(m, int32(_a_F_predicatelock_twophase_recover_0))
							mBase = m.M
							v283 = m.ExcPending
							if v283 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_predicatelock_twophase_recover_1), int32(0))
								mBase = m.M
								v287 = m.ExcPending
								if v287 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_predicatelock_twophase_recover_2), int32(_a_F_predicatelock_twophase_recover_3), int32(_a_F_predicatelock_twophase_recover_4))
									mBase = m.M
									v292 = m.ExcPending
									if v292 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v45))) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = int64(-4294967296)
						*(*int32)(unsafe.Add(mBase, uint32(v23-int32(60)))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v23-int32(48)))) = int64(-1)
						v63 = int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v23-int32(56)))) = v63
						v66 = v23 + int32(24)
						*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v66
						*(*int64)(unsafe.Add(mBase, uint32(v23-int32(40)))) = v63
						*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = l0
						*(*int64)(unsafe.Add(mBase, uint32(v23-int32(8)))) = int64(0)
						v80 = v23 - int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(v23-int32(12)))) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v66))) = v66
						*(*int32)(unsafe.Add(mBase, uint32(v80))) = v80
						v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v84
						v87 = v23 + int32(44)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v87))) = v88
						if v88&int32(32) != 0 {
							v97 = v88
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v92 + int32(1)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
							v97 = v96
						}
						*(*int32)(unsafe.Add(mBase, uint32(v87))) = v97 | int32(1536)
						v104 = v23 - int32(24)
						*(*int32)(unsafe.Add(mBase, uint32(v23-int32(20)))) = v104
						v109 = v23 - int32(32)
						*(*int32)(unsafe.Add(mBase, uint32(v23-int32(28)))) = v109
						*(*int32)(unsafe.Add(mBase, uint32(v104))) = v104
						*(*int32)(unsafe.Add(mBase, uint32(v109))) = v109
						*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l0
						v115 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[2]))
						v121 = F_hash_search(m, v115, v11+int32(12), int32(1), v11+int32(11))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v45
							v125 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
							if v126 != 0 {
								v128 = v23 + int32(40)
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v129))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
									v141 = base.B2i32(base.Ui32(v129) < base.Ui32(v126))
								} else {
									v141 = base.B2i32(int32(0) < v126-v129)
								}
								v143 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
								if v141 == int32(0) {
									v216 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
									v217 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
									if v216 != v217 {
									} else {
										v219 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v143)+20)) = v219 + int32(1)
									}
									v228 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
									F_LWLockRelease(m, v228+int32(3584))
									mBase = m.M
									v232 = m.ExcPending
									if v232 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16)
										return
									}
								} else {
									v147 = v143
									v150 = v23 + int32(40)
									v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
									*(*int32)(unsafe.Add(mBase, uint32(v147)+20)) = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v147)+16)) = v151
									v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
									v157 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
									v161 = F_LWLockAcquire(m, v157+int32(_a_F_predicatelock_twophase_recover_5), int32(0))
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										if v155 == int32(0) {
											v166 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
											*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = int64(0)
										} else {
											v171 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])))
											if v171 == int32(1) {
												v176 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[5]))
												v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+316))
												v179 = base.B2i32(v177 != int32(2))
												*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])) = uint8(v179)
												v181 = v179
											} else {
												v181 = int32(0)
											}
											v183 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
											if v181 == int32(0) {
												v206 = v183
												*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v155
											} else {
												v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
												if v186 == int32(0) {
													v206 = v183
													*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v155
												} else {
													if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v186))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v155)) == int32(0) {
														v200 = base.B2i32(base.Ui32(v155) < base.Ui32(v186))
													} else {
														v200 = int32(base.Ui32(v155-v186) >> (uint(int32(31)) % 32))
													}
													if v200 == int32(0) {
													} else {
														v204 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
														v206 = v204
														*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v155
													}
												}
											}
										}
										v211 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
										F_LWLockRelease(m, v211+int32(_a_F_predicatelock_twophase_recover_5))
										mBase = m.M
										v215 = m.ExcPending
										if v215 != 0 {
											return
										} else {
											v228 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
											F_LWLockRelease(m, v228+int32(3584))
											mBase = m.M
											v232 = m.ExcPending
											if v232 != 0 {
												return
											} else {
												m.G0 = v11 + int32(16)
												return
											}
										}
									}
								}
							} else {
								v147 = v125
								v150 = v23 + int32(40)
								v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
								*(*int32)(unsafe.Add(mBase, uint32(v147)+20)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v147)+16)) = v151
								v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
								v157 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
								v161 = F_LWLockAcquire(m, v157+int32(_a_F_predicatelock_twophase_recover_5), int32(0))
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return
								} else {
									if v155 == int32(0) {
										v166 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
										*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = int64(0)
									} else {
										v171 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])))
										if v171 == int32(1) {
											v176 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[5]))
											v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+316))
											v179 = base.B2i32(v177 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])) = uint8(v179)
											v181 = v179
										} else {
											v181 = int32(0)
										}
										v183 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
										if v181 == int32(0) {
											v206 = v183
											*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v155
										} else {
											v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
											if v186 == int32(0) {
												v206 = v183
												*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v155
											} else {
												if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v186))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v155)) == int32(0) {
													v200 = base.B2i32(base.Ui32(v155) < base.Ui32(v186))
												} else {
													v200 = int32(base.Ui32(v155-v186) >> (uint(int32(31)) % 32))
												}
												if v200 == int32(0) {
												} else {
													v204 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
													v206 = v204
													*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v155
												}
											}
										}
									}
									v211 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
									F_LWLockRelease(m, v211+int32(_a_F_predicatelock_twophase_recover_5))
									mBase = m.M
									v215 = m.ExcPending
									if v215 != 0 {
										return
									} else {
										v228 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
										F_LWLockRelease(m, v228+int32(3584))
										mBase = m.M
										v232 = m.ExcPending
										if v232 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16)
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
	case 1:
		v234 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[6]))
		v236 = l2 + int32(4)
		v237 = F_get_hash_value(m, v234, v236)
		mBase = m.M
		v238 = m.ExcPending
		if v238 != 0 {
			return
		} else {
			v240 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
			v244 = F_LWLockAcquire(m, v240+int32(3584), int32(1))
			mBase = m.M
			v245 = m.ExcPending
			if v245 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
				v248 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[2]))
				v251 = int32(0)
				v253 = F_hash_search(m, v248, v11+int32(4), v251, v251)
				mBase = m.M
				v254 = m.ExcPending
				if v254 != 0 {
					return
				} else {
					v256 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
					F_LWLockRelease(m, v256+int32(3584))
					mBase = m.M
					v260 = m.ExcPending
					if v260 != 0 {
						return
					} else {
						v261 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
						F_CreatePredicateLock(m, v236, v237, v261)
						mBase = m.M
						v263 = m.ExcPending
						if v263 != 0 {
							return
						} else {
							m.G0 = v11 + int32(16)
							return
						}
					}
				}
			}
		}
	default:
		m.G0 = v11 + int32(16)
		return
	}
}
func F_prepare_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_prepare_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(992)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v15
	v19 = int32(_a_F_prepare_cb_wrapper_1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_prepare_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_prepare_cb_wrapper[0])) = v9 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+147)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+164)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if v37 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_prepare_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_prepare_cb_wrapper_3), v9)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_prepare_cb_wrapper_4), int32(986), int32(_a_F_prepare_cb_wrapper_5))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.T0[v37].(func(*base.Module, int32, int32, int64))(m, v11, l1, l2)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_prepare_cb_wrapper[0])) = v60
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_printsimple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v145 int32
	_ = v145
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v358 int32
	_ = v358
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v16 < v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_slot_getsomeattrs_int(m, l0, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_pq_beginmessage(m, v10+int32(-16), int32(68))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	F_enlargeStringInfo(m, v10+int32(-16), int32(2))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v36 = int32(8)
	v40 = v27<<(uint(v36)%32) | int32(base.Ui32(v27)>>(uint(v36)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v33+v34))) = uint16(v40)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v33 + int32(2)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if int32(0) < v45 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v53 = v45
	v54 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_pq_endmessage(m, v10+int32(-16))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L72
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v54))))
	if v62 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L10
L13:
	;
	v343 = v54 + int32(1)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v343 < v344 {
		v53 = v344
		v54 = v343
		goto L11
	} else {
		goto L71
	}
L14:
	;
	F_enlargeStringInfo(m, v10+int32(-16), int32(4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v54<<(uint(int32(2))%32))))
	v88 = v14 + int32(88) + v53<<(uint(int32(4))%32) + v54*int32(100)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	switch v89 - int32(20) {
	case 0:
		goto L21
	default:
		goto L19
	case 3:
		goto L18
	case 5:
		goto L22
	case 6:
		goto L20
	}
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v70+v71))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v70 + int32(4)
	goto L13
L18:
	;
	v316 = v10 + int32(-48)
	if int32(0) <= v82 {
		goto L67
	} else {
		goto L68
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L63
	}
L20:
	;
	v163 = v10 + int32(-48)
	if v82 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L21:
	;
	v137 = v10 + int32(-48)
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	if int64(0) <= v138 {
		goto L40
	} else {
		goto L41
	}
L22:
	;
	v94 = F_pg_detoast_datum_packed(m, v82)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v96 = int32(1)
	v97 = v94 + v96
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v102 = v100 & v96
	if v102 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v103 = v97
	goto L26
L25:
	;
	v103 = v94 + int32(4)
	goto L26
L26:
	;
	if v100 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_pq_sendcountedtext(m, v10+int32(-16), v103, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L38
	}
L28:
	;
	v106 = int32(4)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v108&int32(254) == int32(2) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v121 = int32(1)
	if v102 != 0 {
		v131 = int32(base.Ui32(v100)>>(uint(v121)%32)) - v121
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v117 = v106
	goto L33
L32:
	;
	v117 = base.B2i32(v108 == int32(18)) << (uint(v106) % 32)
	goto L33
L33:
	;
	if v108 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v120 = v106
	goto L36
L35:
	;
	v120 = v117
	goto L36
L36:
	;
	v131 = v120
	goto L27
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v131 = int32(base.Ui32(v125)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	goto L13
L39:
	;
	F_pq_sendcountedtext(m, v10+int32(-16), v137, v154)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L43
	}
L40:
	;
	v150 = v138
	v151 = int32(0)
	goto L42
L41:
	;
	v145 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v145)
	v150 = int64(0) - v138
	v151 = int32(1)
	goto L42
L42:
	;
	v153 = F_pg_ulltoa_n(m, v150, v137+v151)
	mBase = m.M
	v154 = v153 + v151
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137+v154))) = uint8(v156)
	goto L39
L43:
	;
	goto L13
L44:
	;
	F_pq_sendcountedtext(m, v10+int32(-16), v163, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L62
	}
L45:
	;
	v175 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v163))) = uint8(v175)
	v296 = int32(1)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v181 = int32(1233)
	v186 = int32(base.Ui32((base.I32_clz(v82)^int32(31))*v181+v181) >> (uint(int32(12)) % 32))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v186<<(uint(int32(2))%32))+uint32(_c_F_printsimple[0])))
	v193 = v186 + base.B2i32(base.Ui32(v191) <= base.Ui32(v82))
	v194 = int32(0)
	if base.Ui32(v82) < base.Ui32(int32(_a_F_printsimple_0)) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if base.Ui32(v241) < base.Ui32(int32(100)) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v240 = v194
	v241 = v82
	goto L48
L50:
	;
	goto L51
L51:
	;
	v198 = v82
	v200 = v194
	goto L52
L52:
	;
	v207 = v163 + v193 - v200
	v208 = int32(4)
	v211 = base.I32_div_u_s(v198, int32(_a_F_printsimple_0))
	v214 = v211*int32(-10000) + v198
	v215 = int32(100)
	v216 = base.I32_div_u_s(v214, v215)
	v217 = int32(1)
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216<<(uint(v217)%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v207-v208))) = uint16(v221)
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v214-v216*v215)<<(uint(v217)%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v207-int32(2)))) = uint16(v232)
	v235 = v200 + v208
	if base.Ui32(int32(99999999)) < base.Ui32(v198) {
		v198 = v211
		v200 = v235
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v240 = v235
	v241 = v211
	goto L48
L54:
	;
	goto L53
L55:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v270) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v270 = v241
	v271 = v240
	goto L55
L57:
	;
	goto L58
L58:
	;
	v251 = int32(2)
	v253 = int32(_a_F_printsimple_1)
	v255 = int32(100)
	v256 = base.I32_div_u_s(v241&v253, v255)
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v241-v256*v255)&v253<<(uint(int32(1))%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v163+v193-v240-v251))) = uint16(v266)
	v270 = v256
	v271 = v240 | v251
	goto L55
L59:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v270<<(uint(int32(1))%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v163+v193-v271-int32(2)))) = uint16(v282)
	v296 = v193
	goto L44
L60:
	;
	goto L61
L61:
	;
	v285 = v270 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v163))) = uint8(v285)
	v296 = v193
	goto L44
L62:
	;
	goto L13
L63:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v303
	F_errmsg_internal(m, int32(_a_F_printsimple_2), v12)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_printsimple_3), int32(136), int32(_a_F_printsimple_4))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_pq_sendcountedtext(m, v10+int32(-16), v316, v332)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L70
	}
L67:
	;
	v328 = v82
	v329 = int32(0)
	goto L69
L68:
	;
	v323 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v316))) = uint8(v323)
	v328 = int32(0) - v82
	v329 = int32(1)
	goto L69
L69:
	;
	v331 = F_pg_ultoa_n(m, v328, v316+v329)
	mBase = m.M
	v332 = v331 + v329
	v334 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v316+v332))) = uint8(v334)
	goto L66
L70:
	;
	goto L13
L71:
	;
	goto L12
L72:
	;
	m.G0 = v12 - int32(-64)
	return int32(1)
}
func F_privilege_to_string(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v1 = l0
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if v1 <= int64(127) {
		v12 = v1 - int64(1)
		if base.Ui64(int64(63)) < base.Ui64(v12) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v1)
				F_errmsg_internal(m, int32(_a_F_privilege_to_string_0), v7)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_privilege_to_string_1), int32(2643), int32(_a_F_privilege_to_string_2))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			switch base.I32_wrap_i64(v12) - int32(1) {
			case 0:
				v72 = int32(_a_F_privilege_to_string_3)
				m.G0 = v7 + int32(16)
				return v72
			case 1, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v1)
					F_errmsg_internal(m, int32(_a_F_privilege_to_string_0), v7)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_privilege_to_string_1), int32(2643), int32(_a_F_privilege_to_string_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 2:
				v72 = int32(_a_F_privilege_to_string_4)
				m.G0 = v7 + int32(16)
				return v72
			case 6:
				v72 = int32(_a_F_privilege_to_string_5)
				m.G0 = v7 + int32(16)
				return v72
			case 14:
				v72 = int32(_a_F_privilege_to_string_6)
				m.G0 = v7 + int32(16)
				return v72
			case 30:
				v72 = int32(_a_F_privilege_to_string_7)
				m.G0 = v7 + int32(16)
				return v72
			case 62:
				v72 = int32(_a_F_privilege_to_string_8)
				m.G0 = v7 + int32(16)
				return v72
			default:
				v72 = int32(_a_F_privilege_to_string_9)
				m.G0 = v7 + int32(16)
				return v72
			}
		}
	} else {
		if v1 <= int64(2047) {
			if v1 <= int64(511) {
				if v1 == int64(128) {
					v72 = int32(_a_F_privilege_to_string_10)
					m.G0 = v7 + int32(16)
					return v72
				} else {
					if v1 != int64(256) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v1)
							F_errmsg_internal(m, int32(_a_F_privilege_to_string_0), v7)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_privilege_to_string_1), int32(2643), int32(_a_F_privilege_to_string_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v72 = int32(_a_F_privilege_to_string_11)
						m.G0 = v7 + int32(16)
						return v72
					}
				}
			} else {
				if v1 == int64(512) {
					v72 = int32(_a_F_privilege_to_string_12)
					m.G0 = v7 + int32(16)
					return v72
				} else {
					if v1 != int64(1024) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v1)
							F_errmsg_internal(m, int32(_a_F_privilege_to_string_0), v7)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_privilege_to_string_1), int32(2643), int32(_a_F_privilege_to_string_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v72 = int32(_a_F_privilege_to_string_13)
						m.G0 = v7 + int32(16)
						return v72
					}
				}
			}
		} else {
			if v1 <= int64(8191) {
				if v1 == int64(2048) {
					v72 = int32(_a_F_privilege_to_string_14)
					m.G0 = v7 + int32(16)
					return v72
				} else {
					if v1 != int64(4096) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v1)
							F_errmsg_internal(m, int32(_a_F_privilege_to_string_0), v7)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_privilege_to_string_1), int32(2643), int32(_a_F_privilege_to_string_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v72 = int32(_a_F_privilege_to_string_15)
						m.G0 = v7 + int32(16)
						return v72
					}
				}
			} else {
				if v1 == int64(8192) {
					v72 = int32(_a_F_privilege_to_string_16)
					m.G0 = v7 + int32(16)
					return v72
				} else {
					if v1 != int64(16384) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v1)
							F_errmsg_internal(m, int32(_a_F_privilege_to_string_0), v7)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_privilege_to_string_1), int32(2643), int32(_a_F_privilege_to_string_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v72 = int32(_a_F_privilege_to_string_17)
						m.G0 = v7 + int32(16)
						return v72
					}
				}
			}
		}
	}
}
func F_processPendingPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v5
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(-1)
	if base.Ui32(int32(25)) <= base.Ui32(v19) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = int32(base.Ui32(v19+int32(_a_F_processPendingPage_0)) >> (uint(int32(2)) % 32))
	goto L3
L2:
	;
	v31 = v5
	goto L3
L3:
	;
	v33 = v31 & int32(_a_F_processPendingPage_1)
	if base.Ui32(l3) <= base.Ui32(v33) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = l3
	v44 = v5
	goto L7
L5:
	;
	v152 = v5
	v154 = v5
	goto L6
L6:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_ginInsertBAEntries(m, l0, v15+int32(8), v154&int32(_a_F_processPendingPage_1), v163, v164, v152)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L33
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v40&int32(_a_F_processPendingPage_1)<<(uint(int32(2))%32)+(l2+int32(24))-int32(4))))
	v60 = l2 + v57&int32(_a_F_processPendingPage_2)
	v61 = F_gintuple_get_attrnum(m, v49, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v152 = v140
	v154 = v103
	goto L6
L9:
	;
	return
L10:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)))
	if v63 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v107 = F_gintuple_get_key(m, v104, v60, v15+int32(7))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L26
	}
L12:
	;
	v65 = v15 + int32(8)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+2)))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	v68 = int32(16)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)))
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	if v66|v67<<(uint(v68)%32) == v71|v72<<(uint(v68)%32) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	goto L14
L14:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v99)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v101
	v103 = v61
	goto L11
L15:
	;
	if v61 == v44&int32(_a_F_processPendingPage_1) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	goto L15
L17:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)))
	if v78 == v79 {
		v82 = int32(1)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v82 = int32(0)
	goto L16
L20:
	;
	goto L19
L21:
	;
	v87 = v82
	goto L23
L22:
	;
	v87 = int32(0)
	goto L23
L23:
	;
	if v87 != 0 {
		v103 = v44
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_ginInsertBAEntries(m, l0, v15+int32(8), v44&int32(_a_F_processPendingPage_1), v92, v93, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	goto L14
L26:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v111 <= v110 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v111 << (uint(int32(1)) % 32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v119 = F_repalloc(m, v116, v111<<(uint(int32(3))%32))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	v128 = v110
	goto L29
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v128<<(uint(int32(2))%32)))) = v107
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v134+v135))) = uint8(v109)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v139 = int32(1)
	v140 = v138 + v139
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v140
	v143 = v40 + v139
	if base.Ui32(v143&int32(_a_F_processPendingPage_1)) <= base.Ui32(v33) {
		v40 = v143
		v44 = v103
		goto L7
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v119
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v124 = F_repalloc(m, v122, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v124
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v128 = v127
	goto L29
L32:
	;
	goto L8
L33:
	;
	m.G0 = v15 + int32(16)
	return
}
func F_process_shmem_requests(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_shmem_requests[0])) = uint8(v3)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_process_shmem_requests[1]))
	if v6 != 0 {
		m.T0[v6].(func(*base.Module))(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_process_shmem_requests[0])) = uint8(v10)
			return
		}
	} else {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_process_shmem_requests[0])) = uint8(v10)
		return
	}
}
func F_process_syncing_tables(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int64
	_ = v141
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v334 int32
	_ = v334
	var v337 int64
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int64
	_ = v610
	var v611 int64
	_ = v611
	var v619 int64
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int64
	_ = v630
	var v632 int32
	_ = v632
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(160)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	switch v22 {
	case 0:
		goto L3
	case 1:
		goto L5
	case 2:
		goto L4
	default:
		goto L1
	}
L1:
	;
	m.G0 = v18 + int32(160)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+56)) = int32(0)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L9
	} else {
		goto L183
	}
L4:
	;
	v157 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)) = uint8(v157)
	v161 = F_FetchTableStates(m, v18+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L29
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = int32(1)
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	F_s_lock(m, v27+int32(56), int32(_a_F_process_syncing_tables_0), int32(296), int32(_a_F_process_syncing_tables_1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+40)))
	if v37 != int32(99) {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v36)+48))
	if base.Ui64(l0) < base.Ui64(v40) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v42 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+152)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18-int32(-64)))) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v36)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = l0
	v79 = int32(115)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+40)) = uint8(v79)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[1]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	goto L13
L13:
	;
	if base.B2i32(v83 == int32(2)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v94 = int32(*(*int8)(unsafe.Add(mBase, uint32(v91)+40)))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v91)+48))
	F_UpdateSubscriptionRelState(m, v92, v93, v94, v95, int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[2]))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[3]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+36))
	m.T0[v105].(func(*base.Module, int32, int32))(m, v100, v18+int32(92))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+32))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+36))
	F_ReplicationSlotNameForTablesync(m, v110, v111, v18+int32(96))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[2]))
	F_ReplicationSlotDropAtPubNode(m, v117, v18+int32(96), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v126 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+32))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+36))
	F_ReplicationOriginNameForLogicalRep(m, v132, v133, v18+int32(16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_replorigin_session_reset(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v141 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_process_syncing_tables[4])) = v141
	*(*int64)(unsafe.Add(mBase, _c_F_process_syncing_tables[5])) = v141
	v147 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _c_F_process_syncing_tables[6])) = uint16(v147)
	F_replorigin_drop_by_name(m, v18+int32(16), int32(1), v147)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	F_finish_sync_worker(m)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[7]))
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[8]))
	if v166 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[8]))
	if v188 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	if v166 != 0 {
		goto L30
	} else {
		goto L35
	}
L32:
	;
	if v164 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = int64(68719476740)
	v177 = F_hash_create(m, int32(_a_F_process_syncing_tables_2), int32(256), v18+int32(96), int32(40))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[7])) = v177
	goto L30
L35:
	;
	if v164 == int32(0) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	F_hash_destroy(m, v164)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[7])) = int32(0)
	goto L30
L38:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	if v693 != int32(1) {
		goto L1
	} else {
		goto L160
	}
L39:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v191 <= int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v200 = v2
	v205 = v2
	goto L41
L41:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+v205<<(uint(int32(2))%32))))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+16)))
	if v214 == int32(115) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v660 == int32(0) {
		goto L38
	} else {
		goto L158
	}
L43:
	;
	v670 = v205 + int32(1)
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v670 < v671 {
		v200 = v660
		v205 = v670
		goto L41
	} else {
		goto L157
	}
L44:
	;
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v213)+8))
	if base.Ui64(l0) < base.Ui64(v217) {
		v660 = v200
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	v271 = F_LWLockAcquire(m, v267+int32(_a_F_process_syncing_tables_3), int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L9
	} else {
		goto L60
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v213)+8)) = l0
	v220 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+16)) = uint8(v220)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	if v222 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L9
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+32))
	F_LockSharedObject(m, int32(_a_F_process_syncing_tables_4), v232, int32(1))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L9
	} else {
		goto L52
	}
L51:
	;
	v227 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)) = uint8(v227)
	goto L50
L52:
	;
	if v200 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v240 = F_table_open(m, int32(_a_F_process_syncing_tables_5), int32(3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L9
	} else {
		goto L56
	}
L54:
	;
	v242 = v200
	goto L55
L55:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	F_ReplicationOriginNameForLogicalRep(m, v245, v246, v18+int32(96))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L57
	}
L56:
	;
	v242 = v240
	goto L55
L57:
	;
	F_replorigin_drop_by_name(m, v18+int32(96), int32(1), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+32))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v261 = int32(*(*int8)(unsafe.Add(mBase, uint32(v213)+16)))
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v213)+8))
	F_UpdateSubscriptionRelState(m, v259, v260, v261, v262, int32(1))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	v660 = v242
	goto L43
L60:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+32))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v277 = int32(0)
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[10]))
	if v282 <= v277 {
		v314 = v277
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v314 != 0 {
		goto L72
	} else {
		goto L73
	}
L62:
	;
	goto L61
L63:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[11]))
	v292 = v277
	goto L64
L64:
	;
	v297 = v286 + int32(16) + v292*int32(112)
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+16)))
	if v298 != int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v314 = int32(0)
	goto L62
L66:
	;
	v309 = v292 + int32(1)
	if v309 != v282 {
		v292 = v309
		goto L64
	} else {
		goto L71
	}
L67:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	if v301 == int32(3) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v297)+32))
	if v304 != v275 {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v297)+36))
	if v306 != v276 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v314 = v297
	goto L62
L71:
	;
	goto L65
L72:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v314)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v314)+56)) = int32(1)
	v322 = v314 + int32(56)
	if v318 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)+32))
	v484 = int32(0)
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[10]))
	if v487 <= v484 {
		v582 = v484
		goto L130
	} else {
		goto L131
	}
L75:
	;
	F_s_lock(m, v322, int32(_a_F_process_syncing_tables_0), int32(537), int32(_a_F_process_syncing_tables_6))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L9
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+16)) = uint8(v328)
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v314)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v213)+8)) = v330
	if v328 == int32(119) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	v334 = int32(99)
	*(*uint8)(unsafe.Add(mBase, uint32(v314)+40)) = uint8(v334)
	if base.Ui64(l0) < base.Ui64(v330) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = int32(0)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+16)))
	if v341 == int32(119) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v337 = v330
	goto L84
L83:
	;
	v337 = l0
	goto L84
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v314)+48)) = v337
	goto L81
L85:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v314)+20))
	if v344 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	F_LWLockRelease(m, v476+int32(_a_F_process_syncing_tables_3))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L9
	} else {
		goto L129
	}
L88:
	;
	F_logicalrep_worker_wakeup_ptr(m, v314)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L9
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	F_LWLockRelease(m, v348+int32(_a_F_process_syncing_tables_3))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L9
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	if v353 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v200 != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v365 = v200
	goto L95
L95:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L9
	} else {
		goto L102
	}
L96:
	;
	F_sequence_close(m, v200, int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L9
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v361 = int32(0)
	v363 = F_pgstat_report_stat(m, v361)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	v365 = v361
	goto L95
L102:
	;
	v368 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)) = uint8(v368)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	goto L103
L103:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[12]))
	if v387 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L9
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L9
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+32))
	v397 = F_GetSubscriptionRelState(m, v394, v370, v18+int32(96))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	if v397 == int32(0) {
		v660 = v365
		goto L43
	} else {
		goto L111
	}
L111:
	;
	if v397&int32(255) == int32(115) {
		v660 = v365
		goto L43
	} else {
		goto L112
	}
L112:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	v410 = F_LWLockAcquire(m, v406+int32(_a_F_process_syncing_tables_3), int32(1))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L9
	} else {
		goto L113
	}
L113:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+32))
	v415 = int32(0)
	v420 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[10]))
	if v420 <= v415 {
		v452 = v415
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	F_LWLockRelease(m, v457+int32(_a_F_process_syncing_tables_3))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L9
	} else {
		goto L125
	}
L115:
	;
	goto L114
L116:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[11]))
	v430 = v415
	goto L117
L117:
	;
	v435 = v424 + int32(16) + v430*int32(112)
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+16)))
	if v436 != int32(1) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v452 = int32(0)
	goto L115
L119:
	;
	v447 = v430 + int32(1)
	if v447 != v420 {
		v430 = v447
		goto L117
	} else {
		goto L124
	}
L120:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	if v439 == int32(3) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v435)+32))
	if v442 != v414 {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v435)+36))
	if v444 != v370 {
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v452 = v435
	goto L115
L124:
	;
	goto L118
L125:
	;
	if v452 == int32(0) {
		v660 = v365
		goto L43
	} else {
		goto L126
	}
L126:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[13]))
	v469 = F_WaitLatch(m, v465, int32(41), int32(1000), int32(134217760))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L9
	} else {
		goto L127
	}
L127:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v472))) = int32(0)
	goto L128
L128:
	;
	goto L103
L129:
	;
	v660 = v200
	goto L43
L130:
	;
	v594 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	F_LWLockRelease(m, v594+int32(_a_F_process_syncing_tables_3))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L9
	} else {
		goto L147
	}
L131:
	;
	v490 = int32(1)
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[11]))
	v495 = v493 + int32(16)
	if v487 != v490 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v504 = v484
	v505 = v484
	v508 = int32(0)
	goto L135
L133:
	;
	v552 = v484
	v553 = v484
	goto L134
L134:
	;
	if v487&v490 == int32(0) {
		v582 = v553
		goto L130
	} else {
		goto L144
	}
L135:
	;
	v518 = v495 + v504*int32(112)
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+16)))
	if v519 != int32(1) {
		v528 = v505
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v552 = v545
	v553 = v543
	goto L134
L137:
	;
	v529 = int32(1)
	v533 = v495 + (v504|v529)*int32(112)
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+16)))
	if v534 != v529 {
		v543 = v528
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	if v522 != int32(1) {
		v528 = v505
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v518)+32))
	v528 = v505 + base.B2i32(v525 == v483)
	goto L137
L140:
	;
	v544 = int32(2)
	v545 = v504 + v544
	v547 = v508 + v544
	if v547 != v487&int32(2147483646) {
		v504 = v545
		v505 = v543
		v508 = v547
		goto L135
	} else {
		goto L143
	}
L141:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	if v537 != int32(1) {
		v543 = v528
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v533)+32))
	v543 = v528 + base.B2i32(v540 == v483)
	goto L140
L143:
	;
	goto L136
L144:
	;
	v568 = v495 + v552*int32(112)
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568)+16)))
	if v569 != int32(1) {
		v582 = v553
		goto L130
	} else {
		goto L145
	}
L145:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	if v572 != int32(1) {
		v582 = v553
		goto L130
	} else {
		goto L146
	}
L146:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v568)+32))
	v582 = v553 + base.B2i32(v575 == v483)
	goto L130
L147:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[14]))
	if v600 <= v582 {
		v660 = v200
		goto L43
	} else {
		goto L148
	}
L148:
	;
	v605 = m.G0
	v606 = int32(16)
	v607 = v605 - v606
	m.G0 = v607
	F___gettimeofday(m, v607)
	mBase = m.M
	v610 = *(*int64)(unsafe.Add(mBase, uint32(v607)))
	v611 = int64(*(*int32)(unsafe.Add(mBase, uint32(v607)+8)))
	m.G0 = v607 + v606
	v619 = v611 + v610*int64(1000000) - int64(946684800000000)
	goto L149
L149:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[7]))
	v625 = F_hash_search(m, v621, v213, int32(1), v18+int32(96))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L9
	} else {
		goto L150
	}
L150:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+96)))
	if v627 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v630 = *(*int64)(unsafe.Add(mBase, uint32(v625)+8))
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[15]))
	goto L154
L152:
	;
	goto L153
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v625)+8)) = v619
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v643)+24))
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[16]))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v646)+16))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v643)+28))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v652 = F_logicalrep_worker_launch(m, int32(1), v644, v647, v648, v649, v650, int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L9
	} else {
		goto L156
	}
L154:
	;
	if base.B2i32(base.I64_extend_i32_s(v632)*int64(1000) <= v619-v630) == int32(0) {
		v660 = v200
		goto L43
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v660 = v200
	goto L43
L157:
	;
	goto L42
L158:
	;
	F_sequence_close(m, v660, int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L9
	} else {
		goto L159
	}
L159:
	;
	goto L38
L160:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[16]))
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+28)))
	if v698 != int32(112) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L9
	} else {
		goto L181
	}
L162:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L9
	} else {
		goto L163
	}
L163:
	;
	v703 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+96)) = uint8(v703)
	v707 = F_FetchTableStates(m, v18+int32(96))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L9
	} else {
		goto L164
	}
L164:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+96)))
	if v709 == int32(1) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L9
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[8]))
	v719 = int32(0)
	if v707&base.B2i32(v718 == v719) == v719 {
		goto L161
	} else {
		goto L170
	}
L168:
	;
	v715 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L9
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v726 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	if v726 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[16]))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v729)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v730
	F_errmsg(m, int32(_a_F_process_syncing_tables_7), v18)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L9
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L9
	} else {
		goto L177
	}
L175:
	;
	F_errfinish(m, int32(_a_F_process_syncing_tables_0), int32(670), int32(_a_F_process_syncing_tables_6))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L9
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v743 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[16]))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	F_ApplyLauncherForgetWorkerStartTime(m, v747)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L9
	} else {
		goto L179
	}
L179:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	v757 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L9
	} else {
		goto L182
	}
L182:
	;
	goto L1
L183:
	;
	F_errmsg_internal(m, int32(_a_F_process_syncing_tables_8), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L9
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_process_syncing_tables_0), int32(718), int32(_a_F_process_syncing_tables_9))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L9
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_prsd_lextype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v6 = F_palloc(m, int32(288))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = int32(1)
	goto L3
L3:
	;
	v15 = int32(12)
	v17 = v6 + v11*v15
	*(*int32)(unsafe.Add(mBase, uint32(v17-v15))) = v11
	v24 = v11 << (uint(int32(2)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_prsd_lextype[0])))
	v28 = F_pstrdup(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+276)) = int32(0)
	return v6
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17-int32(8)))) = v28
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_prsd_lextype[1])))
	v36 = F_pstrdup(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17-int32(4)))) = v36
	v40 = v11 + int32(1)
	if v40 != int32(24) {
		v11 = v40
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
}
func F_pt_contained_poly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v11 = F_point_inside(m, v2, v8, v4+int32(40))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v11 != int32(0))
		}
	}
}
func F_pull_up_union_leaf_queries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = l0
	goto L1
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v25 != int32(142) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L31
	}
L3:
	;
	goto L2
L4:
	;
	if v25 != int32(63) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_pull_up_union_leaf_queries(m, v124, l1, l2, l3, l4)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L30
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v32 = F_palloc0(m, int32(36))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+12)) = int64(0)
	v36 = v30 + l4
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(322)
	v41 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v45 = v44
	goto L12
L11:
	;
	v45 = v41
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v45
	v49 = F_palloc0(m, v45<<(uint(int32(1))%32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v52 == int32(0) {
		v101 = v41
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v101
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v109 = F_lappend(m, v108, v32)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L27
	}
L15:
	;
	v55 = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v56 <= v55 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v101 = v41
	goto L14
L17:
	;
	goto L18
L18:
	;
	v61 = v55
	v63 = v56
	v67 = v41
	goto L19
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v61<<(uint(int32(2))%32))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+26)))
	if v76 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v101 = v91
	goto L14
L21:
	;
	v79 = F_makeVarFromTargetEntry(m, v36, v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	v89 = v63
	v91 = v67
	goto L23
L23:
	;
	v93 = v61 + int32(1)
	if v93 < v89 {
		v61 = v93
		v63 = v89
		v67 = v91
		goto L19
	} else {
		goto L26
	}
L24:
	;
	v81 = F_lappend(m, v67, v79)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v49-int32(2)+v83<<(uint(int32(1))%32)))) = uint16(v83)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v89 = v88
	v91 = v81
	goto L23
L26:
	;
	goto L20
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v109
	v113 = F_palloc0(m, int32(8))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(63)
	v119 = F_pull_up_subqueries_recurse(m, l1, v113, int32(0), v32)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	m.G0 = v13 + int32(16)
	return
L30:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v15 = v127
	goto L1
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v132
	F_errmsg_internal(m, int32(_a_F_pull_up_union_leaf_queries_0), v13)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_pull_up_union_leaf_queries_1), int32(1756), int32(_a_F_pull_up_union_leaf_queries_2))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pull_varattnos_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = int32(0)
	if l0 == v3 {
		v27 = v3
		return v27
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 == int32(6) {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v9 != v10 {
				v27 = v3
				return v27
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v12 != 0 {
					v27 = v3
					return v27
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
					v17 = F_bms_add_member(m, v13, v14+int32(7))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
						return int32(0)
					}
				}
			}
		} else {
			v25 = F_expression_tree_walker_impl(m, l0, int32(896), l1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = v25
				return v27
			}
		}
	}
}
func F_pushValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	v4 = l3
	v5 = l4
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l2 <= int32(2046) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return
L2:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v127 = v125 - v126
	if int32(_a_F_pushValue_0) <= v127 {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	if l2&v21 != 0 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v52 = l1
	v53 = int32(-1)
	v54 = int32(0)
	goto L16
L5:
	;
	v18 = int32(0)
	if l2 == v18 {
		v124 = v18
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v27 = F_errsave_start(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v21 = int32(1)
	if l2 != v21 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v92 = l1
	v93 = int32(-1)
	goto L3
L10:
	;
	return
L11:
	;
	if v27 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v34
	F_errmsg(m, int32(_a_F_pushValue_1), v14)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	F_errsave_finish(m, v26, int32(_a_F_pushValue_2), int32(588), int32(_a_F_pushValue_3))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v60 = int32(24)
	v63 = int32(2)
	v67 = *(*int32)(unsafe.Add(mBase, uint32((v59^int32(base.Ui32(v53)>>(uint(v60)%32)))<<(uint(v63)%32))+uint32(_c_F_pushValue[0])))
	v68 = int32(8)
	v70 = v67 ^ v53<<(uint(v68)%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32((v58^int32(base.Ui32(v70)>>(uint(v60)%32)))<<(uint(v63)%32))+uint32(_c_F_pushValue[0])))
	v81 = v78 ^ v70<<(uint(v68)%32)
	v83 = v52 + v63
	v85 = v54 + v63
	if v85 != l2&int32(-2) {
		v52 = v83
		v53 = v81
		v54 = v85
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v92 = v83
	v93 = v81
	goto L3
L18:
	;
	goto L17
L19:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v106 = *(*int32)(unsafe.Add(mBase, uint32((v98^int32(base.Ui32(v93)>>(uint(int32(24))%32)))<<(uint(int32(2))%32))+uint32(_c_F_pushValue[0])))
	v110 = v106 ^ v93<<(uint(int32(8))%32)
	goto L21
L20:
	;
	v110 = v93
	goto L21
L21:
	;
	v124 = v110 ^ int32(-1)
	goto L2
L22:
	;
	v170 = l2 + int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v173 = v171 - v172
	v174 = v170 + v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v175 <= v174 {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v131 = F_errsave_start(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v151 = F_palloc0(m, int32(12))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L31
	}
L26:
	;
	if v131 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v138
	F_errmsg(m, int32(_a_F_pushValue_4), v14+int32(16))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errsave_finish(m, v130, int32(_a_F_pushValue_2), int32(555), int32(_a_F_pushValue_5))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	goto L22
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v124
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)) = uint8(v4)
	v156 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v156)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = l2&int32(4095) | v127<<(uint(int32(12))%32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v165 = F_lcons(m, v151, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v165
	goto L22
L33:
	;
	v182 = v172
	v183 = v175
	goto L36
L34:
	;
	v205 = v171
	goto L35
L35:
	;
	if l2 != 0 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v189 = v183 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v189
	v191 = F_repalloc(m, v182, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L38
	}
L37:
	;
	v205 = v194
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v191
	v194 = v191 + v173
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v196 <= v174 {
		v182 = v191
		v183 = v196
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v212 = v211 + l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v212
	v214 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v214)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v216 + int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v170 + v220
	goto L1
L41:
	;
	v209 = F__emscripten_memcpy_bulkmem(m, v205, l1, l2)
	mBase = m.M
	goto L43
L42:
	;
	goto L43
L43:
	;
	goto L40
}
