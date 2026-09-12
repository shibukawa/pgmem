package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_apply_handle_commit_internal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	v6 = *(*int64)(unsafe.Add(mBase, _consts[685]))
	if v6 == int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	goto L17
L2:
	;
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _consts[685]))
	if v14 != int64(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v11 + int32(16)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	goto L13
L6:
	;
	return
L7:
	;
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = *(*int64)(unsafe.Add(mBase, _consts[685]))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v22)
	v25 = int64(base.Ui64(v22) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v25)
	F_errmsg(m, int32(541610), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	*(*int64)(unsafe.Add(mBase, _consts[685])) = int64(0)
	goto L5
L11:
	;
	F_errfinish(m, int32(519893), int32(4938), int32(179835))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v45 == int32(2) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[689])) = uint8(v129)
	return
L17:
	;
	if v54 == int32(2) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	F_clear_subscription_skip_lsn(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L39
	}
L21:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, _consts[686])) = v61
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, _consts[687])) = v64
	F_CommitTransactionCommand(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	goto L23
L23:
	;
	if base.Ui32(int32(1)) < base.Ui32(v70) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v74 = F_EndTransactionBlock(m, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v79 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v82 = *(*int64)(unsafe.Add(mBase, _consts[688]))
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v85 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+16)))
	if v86 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v89 == int32(3) {
		goto L16
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[690]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v94
	v97 = F_palloc(m, int32(24))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97)+16)) = v83
	*(*int64)(unsafe.Add(mBase, uint32(v97)+8)) = v82
	v102 = *(*int32)(unsafe.Add(mBase, _consts[691]))
	if v102 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v109
	v111 = int32(4164024)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v97
	*(*int32)(unsafe.Add(mBase, _consts[692])) = v97
	v118 = *(*int32)(unsafe.Add(mBase, _consts[693]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v118
	goto L16
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[692]))
	v109 = v104
	goto L35
L37:
	;
	goto L38
L38:
	;
	v106 = int32(4164024)
	*(*int32)(unsafe.Add(mBase, _consts[691])) = v106
	v109 = v106
	goto L35
L39:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	goto L16
}
func F_apply_handle_delete_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v14
	F_EvalPlanQualInit(m, v10+int32(44), v13, v5, v5, int32(-1), v5)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_TargetPrivilegesCheck(m, v12, int64(2))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			v34 = F_table_slot_create(m, v12, v28+int32(104))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				if l3 != 0 {
					v36 = F_RelationFindReplTupleByIndex(m, v12, l3, l2, v34)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v36 != 0 {
							v48 = F_GetTupleTransactionInfo(m, v34, v10+int32(24), v10+int32(28), v10+int32(32))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								if v48 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									F_TargetPrivilegesCheck(m, v72, int64(8))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											F_EvalPlanQualEnd(m, v10+int32(44))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												m.G0 = v10 + int32(96)
												return
											}
										}
									}
								} else {
									v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+28)))
									v54 = int32(*(*uint16)(unsafe.Add(mBase, _consts[381])))
									if v52 == v54 {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										F_TargetPrivilegesCheck(m, v72, int64(8))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												F_EvalPlanQualEnd(m, v10+int32(44))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													m.G0 = v10 + int32(96)
													return
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v34
										v58 = v10 + int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v58
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v58
										v67 = F_list_make1_impl(m, int32(1), v10)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_ReportApplyConflict(m, v13, l1, int32(15), int32(4), l2, int32(0), v67)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												F_TargetPrivilegesCheck(m, v72, int64(8))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														F_EvalPlanQualEnd(m, v10+int32(44))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return
														} else {
															m.G0 = v10 + int32(96)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v81 = v10 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v81
							v92 = F_list_make1_impl(m, int32(1), v10+int32(4))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								F_ReportApplyConflict(m, v13, l1, int32(15), int32(5), l2, int32(0), v92)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									F_EvalPlanQualEnd(m, v10+int32(44))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										m.G0 = v10 + int32(96)
										return
									}
								}
							}
						}
					}
				} else {
					v38 = F_RelationFindReplTupleSeq(m, v12, l2, v34)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						if v38 == int32(0) {
							v81 = v10 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v81
							v92 = F_list_make1_impl(m, int32(1), v10+int32(4))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								F_ReportApplyConflict(m, v13, l1, int32(15), int32(5), l2, int32(0), v92)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									F_EvalPlanQualEnd(m, v10+int32(44))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										m.G0 = v10 + int32(96)
										return
									}
								}
							}
						} else {
							v48 = F_GetTupleTransactionInfo(m, v34, v10+int32(24), v10+int32(28), v10+int32(32))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								if v48 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									F_TargetPrivilegesCheck(m, v72, int64(8))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											F_EvalPlanQualEnd(m, v10+int32(44))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												m.G0 = v10 + int32(96)
												return
											}
										}
									}
								} else {
									v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+28)))
									v54 = int32(*(*uint16)(unsafe.Add(mBase, _consts[381])))
									if v52 == v54 {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										F_TargetPrivilegesCheck(m, v72, int64(8))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												F_EvalPlanQualEnd(m, v10+int32(44))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													m.G0 = v10 + int32(96)
													return
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v34
										v58 = v10 + int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v58
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v58
										v67 = F_list_make1_impl(m, int32(1), v10)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_ReportApplyConflict(m, v13, l1, int32(15), int32(4), l2, int32(0), v67)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												F_TargetPrivilegesCheck(m, v72, int64(8))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														F_EvalPlanQualEnd(m, v10+int32(44))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return
														} else {
															m.G0 = v10 + int32(96)
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
	}
}
func F_apply_handle_prepare_internal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = m.G0
	v5 = v3 - int32(208)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[677]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_TwoPhaseTransactionGid(m, v9, v10, v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[75]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
		if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v15)) == int32(0) {
			F_BeginTransactionBlock(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_CommitTransactionCommand(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, _consts[686])) = v25
					v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int64)(unsafe.Add(mBase, _consts[687])) = v28
					v30 = F_PrepareTransactionBlock(m, v5)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						m.G0 = v5 + int32(208)
						return
					}
				}
			}
		} else {
			v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int64)(unsafe.Add(mBase, _consts[686])) = v25
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int64)(unsafe.Add(mBase, _consts[687])) = v28
			v30 = F_PrepareTransactionBlock(m, v5)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				m.G0 = v5 + int32(208)
				return
			}
		}
	}
}
func F_attach_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v5 = F_palloc(m, int32(656))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		v11 = *(*int32)(unsafe.Add(mBase, _consts[181]))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v11
		v18 = F__emscripten_memset_bulkmem(m, v5+int32(28), base.I32_extend8_s(int32(0)), int32(624))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = l0 + int32(2048)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = l0 + int32(1496)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l1
		v31 = F_LWLockAcquire(m, l0+int32(1476), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1460))
			if v33 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(533245), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525786), int32(1364), int32(327377))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+1460)) = v33 + int32(1)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+1468))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+652)) = v56
				F_LWLockRelease(m, v55+int32(1476))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					return v5
				}
			}
		}
	}
}
func F_internal_bpchar_pattern_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	v10 = int32(1)
	v11 = l0 + v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v14 = v12 & v10
	if v12 == v10 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v14 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v17 = int32(4)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v19&int32(254) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v32 = int32(1)
	if v14 != 0 {
		v42 = int32(base.Ui32(v12)>>(uint(v32)%32)) - v32
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v28 = v17
	goto L7
L6:
	;
	v28 = base.B2i32(v19 == int32(18)) << (uint(v17) % 32)
	goto L7
L7:
	;
	if v19 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = v17
	goto L10
L9:
	;
	v31 = v28
	goto L10
L10:
	;
	v42 = v31
	goto L1
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	v45 = v11
	goto L14
L13:
	;
	v45 = l0 + int32(4)
	goto L14
L14:
	;
	v51 = v42
	goto L15
L15:
	;
	if v51 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v68 = int32(1)
	v69 = l1 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v72 = v70 & v68
	if v70 == v68 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	v67 = v42 >> (uint(int32(31)) % 32) & v42
	goto L17
L19:
	;
	goto L20
L20:
	;
	v61 = v51 - int32(1)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v61))))
	if v63 == int32(32) {
		v51 = v61
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v67 = v51
	goto L17
L22:
	;
	if v72 != 0 {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v75 = int32(4)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v77&int32(254) == int32(2) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v90 = int32(1)
	if v72 != 0 {
		v100 = int32(base.Ui32(v70)>>(uint(v90)%32)) - v90
		goto L22
	} else {
		goto L32
	}
L26:
	;
	v86 = v75
	goto L28
L27:
	;
	v86 = base.B2i32(v77 == int32(18)) << (uint(v75) % 32)
	goto L28
L28:
	;
	if v77 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v89 = v75
	goto L31
L30:
	;
	v89 = v86
	goto L31
L31:
	;
	v100 = v89
	goto L22
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v100 = int32(base.Ui32(v94)>>(uint(int32(2))%32)) - int32(4)
	goto L22
L33:
	;
	v103 = v69
	goto L35
L34:
	;
	v103 = l1 + int32(4)
	goto L35
L35:
	;
	v109 = v100
	goto L36
L36:
	;
	if v109 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v126 = int32(1)
	if v12&v126 != 0 {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	goto L37
L39:
	;
	v125 = v100 >> (uint(int32(31)) % 32) & v100
	goto L38
L40:
	;
	goto L41
L41:
	;
	v119 = v109 - int32(1)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+v119))))
	if v121 == int32(32) {
		v109 = v119
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v125 = v109
	goto L38
L43:
	;
	return v204
L44:
	;
	v130 = v126
	goto L46
L45:
	;
	v130 = int32(4)
	goto L46
L46:
	;
	v131 = l0 + v130
	v132 = int32(1)
	if v70&v132 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v136 = v132
	goto L49
L48:
	;
	v136 = int32(4)
	goto L49
L49:
	;
	v137 = l1 + v136
	v138 = base.B2i32(v67 < v125)
	if v67 < v125 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v139 = v67
	goto L52
L51:
	;
	v139 = v125
	goto L52
L52:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v139) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	if v201 != 0 {
		v204 = v201
		goto L43
	} else {
		goto L71
	}
L54:
	;
	v201 = int32(0)
	goto L53
L55:
	;
	v175 = v170
	v176 = v171
	v177 = v172
	goto L65
L56:
	;
	if (v131|v137)&int32(3) != 0 {
		v170 = v131
		v171 = v137
		v172 = v139
		goto L55
	} else {
		goto L59
	}
L57:
	;
	v163 = v131
	v164 = v137
	v165 = v139
	goto L58
L58:
	;
	if v165 == int32(0) {
		goto L54
	} else {
		goto L64
	}
L59:
	;
	v147 = v131
	v148 = v137
	v149 = v139
	goto L60
L60:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v152 != v153 {
		v170 = v147
		v171 = v148
		v172 = v149
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v163 = v158
	v164 = v156
	v165 = v160
	goto L58
L62:
	;
	v155 = int32(4)
	v156 = v148 + v155
	v158 = v147 + v155
	v160 = v149 - v155
	if base.Ui32(int32(3)) < base.Ui32(v160) {
		v147 = v158
		v148 = v156
		v149 = v160
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v170 = v163
	v171 = v164
	v172 = v165
	goto L55
L65:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v180 == v181 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v201 = v180 - v181
	goto L53
L67:
	;
	v183 = int32(1)
	v188 = v177 - v183
	if v188 != 0 {
		v175 = v175 + v183
		v176 = v176 + v183
		v177 = v188
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L66
L70:
	;
	goto L54
L71:
	;
	if v67 < v125 {
		v204 = int32(-1)
		goto L43
	} else {
		goto L72
	}
L72:
	;
	v204 = base.B2i32(v125 < v67)
	goto L43
}
func F_internal_citext_pattern_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v7 = int32(1)
	v8 = l0 + v7
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v13 = v11 & v7
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_pfree(m, v44)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L16
	} else {
		goto L56
	}
L2:
	;
	v14 = v8
	goto L4
L3:
	;
	v14 = l0 + int32(4)
	goto L4
L4:
	;
	if v11 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v44 = F_str_tolower(m, v14, v42, int32(100))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v17 = int32(4)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v19&int32(254) == int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v32 = int32(1)
	if v13 != 0 {
		v42 = int32(base.Ui32(v11)>>(uint(v32)%32)) - v32
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v28 = v17
	goto L11
L10:
	;
	v28 = base.B2i32(v19 == int32(18)) << (uint(v17) % 32)
	goto L11
L11:
	;
	if v19 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = v17
	goto L14
L13:
	;
	v31 = v28
	goto L14
L14:
	;
	v42 = v31
	goto L5
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	return int32(0)
L17:
	;
	v48 = int32(1)
	v49 = l1 + v48
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v54 = v52 & v48
	if v54 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v55 = v49
	goto L20
L19:
	;
	v55 = l1 + int32(4)
	goto L20
L20:
	;
	if v52 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v85 = F_str_tolower(m, v55, v83, int32(100))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L32
	}
L22:
	;
	v58 = int32(4)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v60&int32(254) == int32(2) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v73 = int32(1)
	if v54 != 0 {
		v83 = int32(base.Ui32(v52)>>(uint(v73)%32)) - v73
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v69 = v58
	goto L27
L26:
	;
	v69 = base.B2i32(v60 == int32(18)) << (uint(v58) % 32)
	goto L27
L27:
	;
	if v60 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v72 = v58
	goto L30
L29:
	;
	v72 = v69
	goto L30
L30:
	;
	v83 = v72
	goto L21
L31:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v83 = int32(base.Ui32(v77)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v87 = F_strlen(m, v44)
	mBase = m.M
	v88 = F_strlen(m, v85)
	mBase = m.M
	v89 = base.B2i32(v87 < v88)
	if v87 < v88 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = v87
	goto L35
L34:
	;
	v90 = v88
	goto L35
L35:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if v152 != 0 {
		v155 = v152
		goto L1
	} else {
		goto L54
	}
L37:
	;
	v152 = int32(0)
	goto L36
L38:
	;
	v126 = v121
	v127 = v122
	v128 = v123
	goto L48
L39:
	;
	if (v44|v85)&int32(3) != 0 {
		v121 = v44
		v122 = v85
		v123 = v90
		goto L38
	} else {
		goto L42
	}
L40:
	;
	v114 = v44
	v115 = v85
	v116 = v90
	goto L41
L41:
	;
	if v116 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L42:
	;
	v98 = v44
	v99 = v85
	v100 = v90
	goto L43
L43:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v103 != v104 {
		v121 = v98
		v122 = v99
		v123 = v100
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v114 = v109
	v115 = v107
	v116 = v111
	goto L41
L45:
	;
	v106 = int32(4)
	v107 = v99 + v106
	v109 = v98 + v106
	v111 = v100 - v106
	if base.Ui32(int32(3)) < base.Ui32(v111) {
		v98 = v109
		v99 = v107
		v100 = v111
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v121 = v114
	v122 = v115
	v123 = v116
	goto L38
L48:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v131 == v132 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v152 = v131 - v132
	goto L36
L50:
	;
	v134 = int32(1)
	v139 = v128 - v134
	if v139 != 0 {
		v126 = v126 + v134
		v127 = v127 + v134
		v128 = v139
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	goto L37
L54:
	;
	if v87 < v88 {
		v155 = int32(-1)
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v155 = base.B2i32(v88 < v87)
	goto L1
L56:
	;
	F_pfree(m, v85)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	return v155
}
func F_internal_flush_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = l0 + v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v9 = l0 + v8
	if base.Ui32(v7) < base.Ui32(v9) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = v7
	goto L5
L3:
	;
	goto L4
L4:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v80
	goto L1
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	v19 = F_secure_write(m, v17, v14, v9-v14)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	if base.Ui32(v73) < base.Ui32(v9) {
		v14 = v73
		goto L5
	} else {
		goto L22
	}
L8:
	;
	return int32(0)
L9:
	;
	if v19 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v26 == int32(27) {
		v73 = v14
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[582])) = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v68 + v19
	v73 = v19 + v14
	goto L7
L13:
	;
	if v26 == int32(6) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	if v26 == v32 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v53
	v58 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[583])) = v58
	*(*int32)(unsafe.Add(mBase, _consts[8])) = v58
	return int32(-1)
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[582])) = v26
	v38 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	if v38 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(307012), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(522089), int32(1405), int32(236981))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	goto L6
}
func F_internal_get_result_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
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
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v268 int64
	_ = v268
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
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
	var v536 int32
	_ = v536
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v623 int32
	_ = v623
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v675 int32
	_ = v675
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v876 int32
	_ = v876
	var v900 int32
	_ = v900
	v6 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v32 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_ReleaseCatCache(m, v32)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L5
	} else {
		goto L269
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v852
	v876 = v846
	goto L1
L3:
	;
	v841 = int32(0)
	v842 = int32(3)
	if l4 == v841 {
		v876 = v842
		goto L1
	} else {
		goto L268
	}
L4:
	;
	if v39 <= int32(3830) {
		goto L234
	} else {
		goto L235
	}
L5:
	;
	return int32(0)
L6:
	;
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	v38 = v36 + v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+108))
	v40 = F_build_function_result_tupdesc_t(m, v32)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L5
	} else {
		goto L228
	}
L10:
	;
	if v40 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v39
	goto L14
L13:
	;
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v45 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v715 != int32(2249) {
		goto L223
	} else {
		goto L224
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v38)+128))
	v49 = int32(1)
	v53 = v40 + v45<<(uint(int32(4))%32)
	v55 = v53 + int32(88)
	if v45 == v49 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v45&v49 == int32(0) {
		v247 = v187
		v249 = v200
		v250 = v201
		v251 = v202
		v252 = v203
		v253 = v206
		v254 = v207
		v255 = v208
		v256 = v209
		goto L60
	} else {
		goto L61
	}
L18:
	;
	v58 = int32(0)
	v187 = v58
	v200 = v6
	v201 = v6
	v202 = v6
	v203 = v6
	v206 = v6
	v207 = v6
	v208 = v6
	v209 = v6
	v211 = v58
	goto L17
L19:
	;
	goto L20
L20:
	;
	v64 = int32(0)
	v68 = v64
	v69 = v64
	v73 = v6
	v81 = v6
	v82 = v6
	v83 = v6
	v84 = v6
	v87 = v6
	v88 = v6
	v89 = v6
	v90 = v6
	goto L21
L21:
	;
	v93 = v69 * int32(100)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v55+v93)))
	if v95 <= int32(3830) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v187 = v169
	v200 = v170
	v201 = v171
	v202 = v172
	v203 = v173
	v206 = v174
	v207 = v175
	v208 = v176
	v209 = v177
	v211 = v179 * int32(100)
	goto L17
L23:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v93+(v53+int32(188)))))
	if v137 <= int32(3830) {
		goto L44
	} else {
		goto L45
	}
L24:
	;
	v127 = int32(1)
	v128 = v118
	v129 = v119
	v130 = v120
	v131 = v121
	v132 = v122
	v133 = v123
	v134 = v124
	v135 = v125
	goto L23
L25:
	;
	v118 = v81
	v119 = v82
	v120 = v83
	v121 = int32(1)
	v122 = v87
	v123 = v88
	v124 = v89
	v125 = v90
	goto L24
L26:
	;
	switch v95 - int32(2277) {
	case 0:
		goto L25
	case 1, 2, 3, 4, 5:
		v127 = v68
		v128 = v81
		v129 = v82
		v130 = v83
		v131 = v84
		v132 = v87
		v133 = v88
		v134 = v89
		v135 = v90
		goto L23
	case 6:
		goto L29
	default:
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	switch v95 - int32(5077) {
	case 0, 2:
		goto L36
	case 1:
		goto L35
	case 3:
		goto L34
	default:
		goto L37
	}
L29:
	;
	v118 = int32(1)
	v119 = v82
	v120 = v83
	v121 = v84
	v122 = v87
	v123 = v88
	v124 = v89
	v125 = v90
	goto L24
L30:
	;
	if v95 == int32(2776) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v95 != int32(3500) {
		v127 = v68
		v128 = v81
		v129 = v82
		v130 = v83
		v131 = v84
		v132 = v87
		v133 = v88
		v134 = v89
		v135 = v90
		goto L23
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v118 = v81
	v119 = v82
	v120 = v83
	v121 = v84
	v122 = v87
	v123 = v88
	v124 = v89
	v125 = int32(1)
	goto L24
L34:
	;
	v118 = v81
	v119 = v82
	v120 = int32(1)
	v121 = v84
	v122 = v87
	v123 = v88
	v124 = v89
	v125 = v90
	goto L24
L35:
	;
	v118 = v81
	v119 = v82
	v120 = v83
	v121 = v84
	v122 = v87
	v123 = v88
	v124 = int32(1)
	v125 = v90
	goto L24
L36:
	;
	v118 = v81
	v119 = int32(1)
	v120 = v83
	v121 = v84
	v122 = v87
	v123 = v88
	v124 = v89
	v125 = v90
	goto L24
L37:
	;
	switch v95 - int32(4537) {
	case 0:
		goto L38
	case 1:
		goto L33
	default:
		goto L39
	}
L38:
	;
	v118 = v81
	v119 = v82
	v120 = v83
	v121 = v84
	v122 = v87
	v123 = int32(1)
	v124 = v89
	v125 = v90
	goto L24
L39:
	;
	if v95 != int32(3831) {
		v127 = v68
		v128 = v81
		v129 = v82
		v130 = v83
		v131 = v84
		v132 = v87
		v133 = v88
		v134 = v89
		v135 = v90
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v118 = v81
	v119 = v82
	v120 = v83
	v121 = v84
	v122 = int32(1)
	v123 = v88
	v124 = v89
	v125 = v90
	goto L24
L41:
	;
	v178 = int32(2)
	v179 = v69 + v178
	v181 = v73 + v178
	if v181 != v45&int32(2147483646) {
		v68 = v169
		v69 = v179
		v73 = v181
		v81 = v170
		v82 = v171
		v83 = v172
		v84 = v173
		v87 = v174
		v88 = v175
		v89 = v176
		v90 = v177
		goto L21
	} else {
		goto L59
	}
L42:
	;
	v169 = int32(1)
	v170 = v160
	v171 = v161
	v172 = v162
	v173 = v163
	v174 = v164
	v175 = v165
	v176 = v166
	v177 = v167
	goto L41
L43:
	;
	v160 = v128
	v161 = v129
	v162 = v130
	v163 = int32(1)
	v164 = v132
	v165 = v133
	v166 = v134
	v167 = v135
	goto L42
L44:
	;
	switch v137 - int32(2277) {
	case 0:
		goto L43
	case 1, 2, 3, 4, 5:
		v169 = v127
		v170 = v128
		v171 = v129
		v172 = v130
		v173 = v131
		v174 = v132
		v175 = v133
		v176 = v134
		v177 = v135
		goto L41
	case 6:
		goto L47
	default:
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	switch v137 - int32(5077) {
	case 0, 2:
		goto L52
	case 1:
		goto L53
	case 3:
		goto L54
	default:
		goto L55
	}
L47:
	;
	v160 = int32(1)
	v161 = v129
	v162 = v130
	v163 = v131
	v164 = v132
	v165 = v133
	v166 = v134
	v167 = v135
	goto L42
L48:
	;
	if v137 == int32(2776) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if v137 != int32(3500) {
		v169 = v127
		v170 = v128
		v171 = v129
		v172 = v130
		v173 = v131
		v174 = v132
		v175 = v133
		v176 = v134
		v177 = v135
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v160 = v128
	v161 = v129
	v162 = v130
	v163 = v131
	v164 = v132
	v165 = int32(1)
	v166 = v134
	v167 = v135
	goto L42
L52:
	;
	v160 = v128
	v161 = int32(1)
	v162 = v130
	v163 = v131
	v164 = v132
	v165 = v133
	v166 = v134
	v167 = v135
	goto L42
L53:
	;
	v160 = v128
	v161 = v129
	v162 = v130
	v163 = v131
	v164 = v132
	v165 = v133
	v166 = int32(1)
	v167 = v135
	goto L42
L54:
	;
	v160 = v128
	v161 = v129
	v162 = int32(1)
	v163 = v131
	v164 = v132
	v165 = v133
	v166 = v134
	v167 = v135
	goto L42
L55:
	;
	switch v137 - int32(4537) {
	case 0:
		goto L51
	case 1:
		goto L56
	default:
		goto L57
	}
L56:
	;
	v160 = v128
	v161 = v129
	v162 = v130
	v163 = v131
	v164 = v132
	v165 = v133
	v166 = v134
	v167 = int32(1)
	goto L42
L57:
	;
	if v137 != int32(3831) {
		v169 = v127
		v170 = v128
		v171 = v129
		v172 = v130
		v173 = v131
		v174 = v132
		v175 = v133
		v176 = v134
		v177 = v135
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v160 = v128
	v161 = v129
	v162 = v130
	v163 = v131
	v164 = int32(1)
	v165 = v133
	v166 = v134
	v167 = v135
	goto L42
L59:
	;
	goto L22
L60:
	;
	if v247&int32(1) == int32(0) {
		goto L15
	} else {
		goto L79
	}
L61:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+v55)))
	if v215 <= int32(3830) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v247 = int32(1)
	v249 = v238
	v250 = v239
	v251 = v240
	v252 = v241
	v253 = v242
	v254 = v243
	v255 = v244
	v256 = v245
	goto L60
L63:
	;
	v238 = v200
	v239 = v201
	v240 = v202
	v241 = int32(1)
	v242 = v206
	v243 = v207
	v244 = v208
	v245 = v209
	goto L62
L64:
	;
	switch v215 - int32(2277) {
	case 0:
		goto L63
	case 1, 2, 3, 4, 5:
		v247 = v187
		v249 = v200
		v250 = v201
		v251 = v202
		v252 = v203
		v253 = v206
		v254 = v207
		v255 = v208
		v256 = v209
		goto L60
	case 6:
		goto L67
	default:
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	switch v215 - int32(5077) {
	case 0, 2:
		goto L72
	case 1:
		goto L73
	case 3:
		goto L74
	default:
		goto L75
	}
L67:
	;
	v238 = int32(1)
	v239 = v201
	v240 = v202
	v241 = v203
	v242 = v206
	v243 = v207
	v244 = v208
	v245 = v209
	goto L62
L68:
	;
	if v215 == int32(2776) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	if v215 != int32(3500) {
		v247 = v187
		v249 = v200
		v250 = v201
		v251 = v202
		v252 = v203
		v253 = v206
		v254 = v207
		v255 = v208
		v256 = v209
		goto L60
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v238 = v200
	v239 = v201
	v240 = v202
	v241 = v203
	v242 = v206
	v243 = int32(1)
	v244 = v208
	v245 = v209
	goto L62
L72:
	;
	v238 = v200
	v239 = int32(1)
	v240 = v202
	v241 = v203
	v242 = v206
	v243 = v207
	v244 = v208
	v245 = v209
	goto L62
L73:
	;
	v238 = v200
	v239 = v201
	v240 = v202
	v241 = v203
	v242 = v206
	v243 = v207
	v244 = int32(1)
	v245 = v209
	goto L62
L74:
	;
	v238 = v200
	v239 = v201
	v240 = int32(1)
	v241 = v203
	v242 = v206
	v243 = v207
	v244 = v208
	v245 = v209
	goto L62
L75:
	;
	switch v215 - int32(4537) {
	case 0:
		goto L71
	case 1:
		goto L76
	default:
		goto L77
	}
L76:
	;
	v238 = v200
	v239 = v201
	v240 = v202
	v241 = v203
	v242 = v206
	v243 = v207
	v244 = v208
	v245 = int32(1)
	goto L62
L77:
	;
	if v215 != int32(3831) {
		v247 = v187
		v249 = v200
		v250 = v201
		v251 = v202
		v252 = v203
		v253 = v206
		v254 = v207
		v255 = v208
		v256 = v209
		goto L60
	} else {
		goto L78
	}
L78:
	;
	v238 = v200
	v239 = v201
	v240 = v202
	v241 = v203
	v242 = int32(1)
	v243 = v207
	v244 = v208
	v245 = v209
	goto L62
L79:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	v263 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v263
	v268 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+52)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v29)+36)) = v268
	if v263 < v48 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v279 = int32(0)
	v286 = v279
	v290 = v263
	v291 = v263
	v292 = v263
	v293 = v263
	v294 = v279
	v295 = v279
	v297 = v279
	v308 = v6
	goto L84
L82:
	;
	v381 = v263
	v382 = v263
	v383 = v263
	v384 = v263
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v381
	if base.B2i32(v381 == int32(0))&v249 != 0 {
		goto L128
	} else {
		goto L129
	}
L84:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(136)+v286<<(uint(int32(2))%32))))
	if v312 <= int32(3830) {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v358
	v381 = v357
	v382 = v358
	v383 = v359
	v384 = v360
	goto L83
L86:
	;
	v366 = v286 + int32(1)
	if v366 != v48 {
		v286 = v366
		v290 = v357
		v291 = v358
		v292 = v359
		v293 = v360
		v294 = v361
		v295 = v362
		v297 = v363
		v308 = v364
		goto L84
	} else {
		goto L127
	}
L87:
	;
	if v297 != 0 {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L124
	}
L88:
	;
	if v294 != 0 {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L121
	}
L89:
	;
	if v293 != 0 {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L118
	}
L90:
	;
	if v292 != 0 {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L115
	}
L91:
	;
	if v295 != 0 {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L112
	}
L92:
	;
	if v291 != 0 {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L109
	}
L93:
	;
	switch v312 - int32(2277) {
	case 0:
		goto L92
	case 1, 2, 3, 4, 5:
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	case 6:
		goto L96
	default:
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	switch v312 - int32(5077) {
	case 0, 2:
		goto L90
	case 1:
		goto L89
	case 3:
		goto L88
	default:
		goto L103
	}
L96:
	;
	if v290 != 0 {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L100
	}
L97:
	;
	if v312 == int32(2776) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	if v312 != int32(3500) {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	v321 = F_get_call_expr_argtype(m, l1, v286)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v321 == int32(0) {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	v357 = v321
	v358 = v291
	v359 = v292
	v360 = v293
	v361 = v294
	v362 = v295
	v363 = v297
	v364 = v308
	goto L86
L103:
	;
	switch v312 - int32(4537) {
	case 0:
		goto L91
	case 1:
		goto L87
	default:
		goto L104
	}
L104:
	;
	if v312 != int32(3831) {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L105
	}
L105:
	;
	if v308 != 0 {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v294
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L106
	}
L106:
	;
	v331 = F_get_call_expr_argtype(m, l1, v286)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	if v331 == int32(0) {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v357 = v290
	v358 = v291
	v359 = v292
	v360 = v293
	v361 = v294
	v362 = v295
	v363 = v297
	v364 = v331
	goto L86
L109:
	;
	v335 = F_get_call_expr_argtype(m, l1, v286)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	if v335 == int32(0) {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	v357 = v290
	v358 = v335
	v359 = v292
	v360 = v293
	v361 = v294
	v362 = v295
	v363 = v297
	v364 = v308
	goto L86
L112:
	;
	v339 = F_get_call_expr_argtype(m, l1, v286)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	if v339 == int32(0) {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	v357 = v290
	v358 = v291
	v359 = v292
	v360 = v293
	v361 = v294
	v362 = v339
	v363 = v297
	v364 = v308
	goto L86
L115:
	;
	v343 = F_get_call_expr_argtype(m, l1, v286)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	if v343 == int32(0) {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	v357 = v290
	v358 = v291
	v359 = v343
	v360 = v293
	v361 = v294
	v362 = v295
	v363 = v297
	v364 = v308
	goto L86
L118:
	;
	v347 = F_get_call_expr_argtype(m, l1, v286)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	if v347 == int32(0) {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	v357 = v290
	v358 = v291
	v359 = v292
	v360 = v347
	v361 = v294
	v362 = v295
	v363 = v297
	v364 = v308
	goto L86
L121:
	;
	v351 = F_get_call_expr_argtype(m, l1, v286)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	if v351 != 0 {
		v357 = v290
		v358 = v291
		v359 = v292
		v360 = v293
		v361 = v351
		v362 = v295
		v363 = v297
		v364 = v308
		goto L86
	} else {
		goto L123
	}
L123:
	;
	goto L3
L124:
	;
	v353 = F_get_call_expr_argtype(m, l1, v286)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	if v353 == int32(0) {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	v357 = v290
	v358 = v291
	v359 = v292
	v360 = v293
	v361 = v294
	v362 = v295
	v363 = v353
	v364 = v308
	goto L86
L127:
	;
	goto L85
L128:
	;
	F_resolve_anyelement_from_others(m, v27+int32(-16))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L131
	}
L129:
	;
	v410 = v382
	goto L130
L130:
	;
	if base.B2i32(v410 == int32(0))&v252 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v410 = v409
	goto L130
L132:
	;
	F_resolve_anyarray_from_others(m, v27+int32(-16))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	if base.B2i32(v418 == int32(0))&v253 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L134
L136:
	;
	F_resolve_anyrange_from_others(m, v27+int32(-16))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	if base.B2i32(v426 == int32(0))&v254 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L138
L140:
	;
	F_resolve_anymultirange_from_others(m, v27+int32(-16))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L5
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if base.B2i32(v383 == int32(0))&v250 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	goto L142
L144:
	;
	F_resolve_anyelement_from_others(m, v27+int32(-32))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L5
	} else {
		goto L147
	}
L145:
	;
	v442 = v384
	goto L146
L146:
	;
	if base.B2i32(v442 == int32(0))&v255 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v442 = v441
	goto L146
L148:
	;
	F_resolve_anyarray_from_others(m, v27+int32(-32))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	if base.B2i32(v450 == int32(0))&v251 != 0 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L150
L152:
	;
	F_resolve_anyrange_from_others(m, v27+int32(-32))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	if base.B2i32(v458 == int32(0))&v256 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L154
L156:
	;
	F_resolve_anymultirange_from_others(m, v27+int32(-32))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	if v466 != 0 {
		v469 = v466
		goto L161
	} else {
		goto L162
	}
L159:
	;
	goto L158
L160:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v474 != 0 {
		v477 = v474
		goto L166
	} else {
		goto L167
	}
L161:
	;
	v470 = F_get_typcollation(m, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L5
	} else {
		goto L164
	}
L162:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	if v467 != 0 {
		v469 = v467
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v473 = int32(0)
	goto L160
L164:
	;
	v473 = v470
	goto L160
L165:
	;
	v482 = int32(0)
	if v481|v473 == v482 {
		v522 = v482
		v523 = v482
		goto L170
	} else {
		goto L171
	}
L166:
	;
	v478 = F_get_typcollation(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L5
	} else {
		goto L169
	}
L167:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v475 != 0 {
		v477 = v475
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v481 = int32(0)
	goto L165
L169:
	;
	v481 = v478
	goto L165
L170:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v536 = int32(0)
	goto L191
L171:
	;
	v487 = int32(0)
	if l1 == v487 {
		v513 = v487
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v513 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L173:
	;
	goto L172
L174:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v493 = v491 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v493) {
		v513 = v487
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v497 = int32(1) << (uint(v493) % 32)
	if v497&int32(3904) == int32(0) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l1+v509)))
	v513 = v511
	goto L173
L177:
	;
	if v497&int32(5) != 0 {
		v509 = int32(16)
		goto L176
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v509 = int32(24)
	goto L176
L180:
	;
	if v493 != int32(30) {
		v513 = v487
		goto L173
	} else {
		goto L181
	}
L181:
	;
	v509 = int32(12)
	goto L176
L182:
	;
	v522 = v473
	v523 = v481
	goto L170
L183:
	;
	goto L184
L184:
	;
	if v473 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v518 = v513
	goto L187
L186:
	;
	v518 = int32(0)
	goto L187
L187:
	;
	if v481 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v520 = v513
	goto L190
L189:
	;
	v520 = int32(0)
	goto L190
L190:
	;
	v522 = v518
	v523 = v520
	goto L170
L191:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v565 = v40 + int32(20) + v559<<(uint(int32(4))%32) + v536*int32(100)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+68))
	if v566 <= int32(3830) {
		goto L201
	} else {
		goto L202
	}
L192:
	;
	goto L15
L193:
	;
	v687 = v536 + int32(1)
	if v687 != v45 {
		v536 = v687
		goto L191
	} else {
		goto L222
	}
L194:
	;
	F_TupleDescInitEntry(m, v40, base.I32_extend16_s(v536+int32(1)), v565+int32(4), v531, int32(-1), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L5
	} else {
		goto L221
	}
L195:
	;
	F_TupleDescInitEntry(m, v40, base.I32_extend16_s(v536+int32(1)), v565+int32(4), v530, int32(-1), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L5
	} else {
		goto L220
	}
L196:
	;
	v652 = base.I32_extend16_s(v536 + int32(1))
	F_TupleDescInitEntry(m, v40, v652, v565+int32(4), v529, int32(-1), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L5
	} else {
		goto L218
	}
L197:
	;
	v635 = base.I32_extend16_s(v536 + int32(1))
	F_TupleDescInitEntry(m, v40, v635, v565+int32(4), v474, int32(-1), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L5
	} else {
		goto L216
	}
L198:
	;
	F_TupleDescInitEntry(m, v40, base.I32_extend16_s(v536+int32(1)), v565+int32(4), v528, int32(-1), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L5
	} else {
		goto L215
	}
L199:
	;
	if v566 != int32(3831) {
		goto L193
	} else {
		goto L213
	}
L200:
	;
	v598 = base.I32_extend16_s(v536 + int32(1))
	F_TupleDescInitEntry(m, v40, v598, v565+int32(4), v526, int32(-1), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L5
	} else {
		goto L211
	}
L201:
	;
	switch v566 - int32(2277) {
	case 0:
		goto L200
	case 1, 2, 3, 4, 5:
		goto L193
	case 6:
		goto L204
	default:
		goto L205
	}
L202:
	;
	goto L203
L203:
	;
	switch v566 - int32(5077) {
	case 0, 2:
		goto L197
	case 1:
		goto L196
	case 3:
		goto L195
	default:
		goto L210
	}
L204:
	;
	v577 = base.I32_extend16_s(v536 + int32(1))
	F_TupleDescInitEntry(m, v40, v577, v565+int32(4), v466, int32(-1), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L5
	} else {
		goto L208
	}
L205:
	;
	if v566 == int32(2776) {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	if v566 != int32(3500) {
		goto L193
	} else {
		goto L207
	}
L207:
	;
	goto L204
L208:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v584<<(uint(int32(4))%32)+v577*int32(100))+16)) = v522
	goto L209
L209:
	;
	goto L193
L210:
	;
	switch v566 - int32(4537) {
	case 0:
		goto L198
	case 1:
		goto L194
	default:
		goto L199
	}
L211:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v605<<(uint(int32(4))%32)+v598*int32(100))+16)) = v522
	goto L212
L212:
	;
	goto L193
L213:
	;
	F_TupleDescInitEntry(m, v40, base.I32_extend16_s(v536+int32(1)), v565+int32(4), v527, int32(-1), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L5
	} else {
		goto L214
	}
L214:
	;
	goto L193
L215:
	;
	goto L193
L216:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v642<<(uint(int32(4))%32)+v635*int32(100))+16)) = v523
	goto L217
L217:
	;
	goto L193
L218:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v659<<(uint(int32(4))%32)+v652*int32(100))+16)) = v523
	goto L219
L219:
	;
	goto L193
L220:
	;
	goto L193
L221:
	;
	goto L193
L222:
	;
	goto L192
L223:
	;
	v723 = int32(1)
	if l4 != 0 {
		v846 = v723
		v852 = v40
		goto L2
	} else {
		goto L227
	}
L224:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if int32(0) <= v718 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	F_assign_record_type_typmod(m, v40)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L5
	} else {
		goto L226
	}
L226:
	;
	goto L223
L227:
	;
	v876 = v723
	goto L1
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = l0
	F_errmsg_internal(m, int32(48455), v29)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L5
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(522690), int32(446), int32(384616))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L5
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L5
	} else {
		goto L263
	}
L232:
	;
	if l3 != 0 {
		goto L245
	} else {
		goto L246
	}
L233:
	;
	v757 = F_exprType(m, l1)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L5
	} else {
		goto L243
	}
L234:
	;
	switch v39 - int32(2277) {
	case 0, 6:
		goto L233
	case 1, 2, 3, 4, 5:
		v761 = v39
		goto L232
	default:
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if base.Ui32(v39-int32(5077)) < base.Ui32(int32(4)) {
		goto L233
	} else {
		goto L240
	}
L237:
	;
	if v39 == int32(2776) {
		goto L233
	} else {
		goto L238
	}
L238:
	;
	if v39 == int32(3500) {
		goto L233
	} else {
		goto L239
	}
L239:
	;
	v761 = v39
	goto L232
L240:
	;
	if base.Ui32(v39-int32(4537)) < base.Ui32(int32(2)) {
		goto L233
	} else {
		goto L241
	}
L241:
	;
	if v39 == int32(3831) {
		goto L233
	} else {
		goto L242
	}
L242:
	;
	v761 = v39
	goto L232
L243:
	;
	if v757 == int32(0) {
		goto L231
	} else {
		goto L244
	}
L244:
	;
	v761 = v757
	goto L232
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v761
	goto L247
L246:
	;
	goto L247
L247:
	;
	if l4 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L250
L249:
	;
	goto L250
L250:
	;
	v767 = F_get_type_func_class(m, v761, v27+int32(-16))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L5
	} else {
		goto L253
	}
L251:
	;
	v777 = int32(3)
	if l2 == int32(0) {
		v876 = v777
		goto L1
	} else {
		goto L256
	}
L252:
	;
	if l4 == int32(0) {
		v876 = v767
		goto L1
	} else {
		goto L254
	}
L253:
	;
	switch v767 - int32(1) {
	case 0, 1:
		goto L252
	case 2:
		goto L251
	default:
		v876 = v767
		goto L1
	}
L254:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v775 = F_lookup_rowtype_tupdesc_copy(m, v773, int32(-1))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	v846 = v767
	v852 = v775
	goto L2
L256:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v780 != int32(383) {
		v876 = v777
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v783 = int32(1)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v786 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v787 = v783
	goto L260
L259:
	;
	v787 = int32(3)
	goto L260
L260:
	;
	if l4 == int32(0) {
		v876 = v787
		goto L1
	} else {
		goto L261
	}
L261:
	;
	if v786 == int32(0) {
		v876 = v787
		goto L1
	} else {
		goto L262
	}
L262:
	;
	v846 = v783
	v852 = v786
	goto L2
L263:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L5
	} else {
		goto L264
	}
L264:
	;
	v799 = F_format_type_be(m, v39)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L5
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v799
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38 + int32(4)
	F_errmsg(m, int32(202143), v27+int32(-48))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L5
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(522690), int32(498), int32(384616))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L5
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	v846 = v842
	v852 = v841
	goto L2
L269:
	;
	m.G0 = v29 - int32(-64)
	return v876
}
func F_internal_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(327875)
			F_errmsg(m, int32(203170), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(518375), int32(373), int32(293059))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
